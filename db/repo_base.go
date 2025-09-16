package db



type RepoBase interface{
	GetAll[T any]()([]*T, error)
	
}
