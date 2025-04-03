package infrastructure

import (
	"database/sql"

	"github.com/BryanChanona/backend_multi/src/CustomRhythm/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}


func (sql *MySQL)RegisterCustomRhythm(customRhythm domain.CustomRhythmModel) error{

	query := `INSERT INTO ritmopersonalizado (id_user, media_bpm_baja, media_bpm_alta) VALUES (?, ?, ?)`
	_, err := sql.db.Exec(query, customRhythm.Id_user, customRhythm.MediaBpmBaja, customRhythm.MediaBpmAlta)
	if err != nil {
		return err
	}
	return nil
}