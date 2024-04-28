package models

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

type UserParams struct {
	Id     int
	Offset int
	Limit  int
}
