package requestmodels

type PatchRequest struct {
	Name      *string `json:"name"`
	Surname   *string `json:"Surname"`
	BirthYear *int    `json:"birthYear"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
