package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// bodyWriter ใช้สำหรับดักจับ Response Body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

type StandardResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"` // omitempty จะไม่แสดง field นี้ถ้ามัน فاضي
}

// GlobalExceptionHandlerAndLogger คือ Middleware ที่เราต้องการ
func GlobalExceptionHandlerAndLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// --- 1. ส่วนดักจับและ Log Request Body (ขาเข้า) ---
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		// สร้าง Body ใหม่เพื่อให้ Handler อื่นๆ ยังอ่านได้
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))

		logger.Info("Request Received",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"body", string(requestBodyBytes),
		)

		// --- 2. ส่วนดักจับ Response และจัดการ Error ---
		writer := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		// defer/recover() ทำหน้าที่เหมือน try-catch สำหรับดักจับ panic
		defer func() {
			if r := recover(); r != nil {
				// เกิด Panic ที่ไม่คาดคิด (เหมือน catch Exception e)
				logger.Error("Unhandled Panic Recovered", "error", r)

				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"status":  "error",
					"message": "An internal server error occurred",
				})
			}
		}()

		c.Next() // เรียก Handler ตัวหลักให้ทำงาน

		// --- 3. ส่วนจัดการ Response และ Error ที่คาดการณ์ไว้ (ขาออก) ---
		duration := time.Since(startTime)

		// ตรวจสอบว่ามี Error ที่ถูก handler ส่งมาใน context หรือไม่
		if len(c.Errors) > 0 {
			// จัดการ Error ที่คาดการณ์ไว้ (เหมือน catch GeneralNotFoundException)
			err := c.Errors.Last()
			logger.Error("Handled Error", "error", err.Error(), "duration", duration)
			// เราสามารถสร้าง response error ที่นี่ได้ (ในตัวอย่างนี้ handler จัดการเอง)
			// ตัวอย่าง: c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return // จบการทำงาน ไม่ต้อง log response ปกติ
		}

		responseBody := writer.body.Bytes()

		statusCode := writer.Status()
		response := StandardResponse{
			Code:   statusCode,
			Status: "success",
		}

		// พยายาม Unmarshal body เดิมเข้าไปใน Data
		var data interface{}
		if err := json.Unmarshal(responseBody, &data); err != nil {
			// ถ้า body เดิมไม่ใช่ JSON (เช่น เป็น text ธรรมดา) ก็ให้เก็บเป็น string
			response.Data = string(responseBody)
		} else {
			response.Data = data
		}

		// Log และส่ง Response ที่ wrap แล้วกลับไป
		logger.Info("Response Sent", "status_code", statusCode, "duration", duration)
		c.JSON(statusCode, response)

	}
}
