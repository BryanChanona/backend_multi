package UseCase

import "github.com/BryanChanona/backend_multi/src/CustomRhythm/domain"

type RegisterCustomRhythmUC struct {
	db domain.CustomRhythmRepository
}

func NewRegisterCustomRhythmUC(db domain.CustomRhythmRepository) *RegisterCustomRhythmUC {
	return &RegisterCustomRhythmUC{
		db: db,
	}
}

func (uc *RegisterCustomRhythmUC) Execute(custom domain.CustomRhythmModel) error {
	return uc.db.RegisterCustomRhythm(custom)
}