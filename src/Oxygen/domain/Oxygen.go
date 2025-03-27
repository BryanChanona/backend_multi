package domain 

type OxygenModel struct{
	Id_oxygen int `json:"id_oxigeno,omitempty"`
	Id_user int `json:"id_user"`
	Date              string  `json:"date"`
	Time              string  `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
}

type UserOxygen struct {
	Id_oxygen int `json:"id_oxigeno,omitempty"`
	Id_user int `json:"id_user"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
	Date              string  `json:"date"`
	Time              string  `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
}
