package domain


type CustomRhythmRepository interface {
	RegisterCustomRhythm( CustomRhythmModel) error
	UpdateCustomRhythm( idUser int ,custom CustomRhythmModel) error
}