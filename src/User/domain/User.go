package domain

type User struct {
	Id_usuario int `json:"id_usuario,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
	Supervisors []int `json:"supervisors,omitempty"`
}