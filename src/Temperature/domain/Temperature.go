package domain

type Temperature struct {
	Id_temp int `json:"id_temp,omitempty"`
	Id_user  int  `json:"id_user"`
	Date string `json:"date"`
	Time string `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
}
type UserTemperature struct {
	Id_temp           int     `json:"id_temp,omitempty"`
	Id_user           int     `json:"id_user"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	Premium           bool    `json:"premium"`
	Date              string  `json:"date"`
	Time              string  `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
	Password string `json:"password"`
}

