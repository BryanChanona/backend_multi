package domain

type HeartRate struct {
	Id_bpm int `json:"id_bpm,omitempty"`
	Id_user int `json:"id_user"`
	Date string `json:"date"`
	Time string `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
	Id_device int `json:"id_device"`
}

type UserHeartRate struct{
	Id_bpm int `json:"id_bpm,omitempty"`
	Id_user int `json:"id_user"`
	Date string `json:"date"`
	Time string `json:"time"`
	RegisteredMeasure float64 `json:"registeredMeasure"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Premium bool `json:"premium"`
}