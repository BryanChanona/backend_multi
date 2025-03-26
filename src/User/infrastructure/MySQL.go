package infrastructure

import (
	"database/sql"
	"log"

	"github.com/BryanChanona/backend_multi/src/User/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) RegisterUser(user domain.User) error {
	query, err := sql.db.Prepare("INSERT INTO `usuario` (nombre,correo,password,premium) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}
	defer query.Exec()

	_, err = query.Exec(user.Name, user.Email, user.Password, user.Premium)

	if err != nil {
		log.Println("Error saving the user:", err)
		return err
	}
	return nil
}