package stores

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price int
}
