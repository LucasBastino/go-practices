package models

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

type UserParams struct {
	From int `schema:"from"`
	To   int `schema:"To"`
}
