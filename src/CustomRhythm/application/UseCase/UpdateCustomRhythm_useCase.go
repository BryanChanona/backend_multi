package UseCase

import "github.com/BryanChanona/backend_multi/src/CustomRhythm/domain"

type UpdateCustomRhythmUC struct {
	db domain.CustomRhythmRepository
}

func NewUpdateCustomRhythmUC(db domain.CustomRhythmRepository) *UpdateCustomRhythmUC {
	return &UpdateCustomRhythmUC{
		db: db,
	}
}


func (uc *UpdateCustomRhythmUC)UpdateCustomRhythm(idUser int, customRhythm domain.CustomRhythmModel) error{
	return uc.db.UpdateCustomRhythm(idUser, customRhythm)
}