package main

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(todo *Todo) (*Todo, error) {
	q := "INSERT INTO Todos (title, completed) VALUSE($1, $2) RETURNING id"
	err := r.db.QueryRow(q, todo.Title, todo.Completed).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *Repo) GetAll() ([]*Todo, error) {
	q := "SELECT id, title, completed FROM todos"
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
