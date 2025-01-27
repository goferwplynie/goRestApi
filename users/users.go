package users

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthYear int    `json:"birthYear"`
}
