package entity

type CreateUserInput struct {
	Email     string
	FirstName string
	LastName  string
	Gender    string
}
type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	Gender    string
}
