package domain


type CustomRhythmModel struct {
	Id_custom_rhythm int `json:"idCustom,omitempty"`
	Id_user int `json:"id_user"`
	MediaBpmBaja string `json:"mediaBaja"`
	MediaBpmAlta string `json:"mediaAlta"`
}