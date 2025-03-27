package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/BryanChanona/backend_multi/src/Oxygen/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) SaveOxygen(oxygen domain.OxygenModel) error {

	if oxygen.Date == "" {
        oxygen.Date = time.Now().Format("2006-01-02")  // Formato YYYY-MM-DD
    }

    // Si no se proporciona hora, usa la hora actual
    if oxygen.Time == "" {
        oxygen.Time = time.Now().Format("15:04")  // Formato HH:mm
    }

    // Parsear la fecha
    fecha, err := time.Parse("2006-01-02", oxygen.Date)
    if err != nil {
        return fmt.Errorf("formato de fecha inválido: %v", err)
    }

    // Parsear la hora
    hora, err := time.Parse("15:04", oxygen.Time)
    if err != nil {
        return fmt.Errorf("formato de hora inválido: %v", err)
    }

	query, err := sql.db.Prepare("INSERT INTO registrooxigeno (id_user, fecha, hora, medidaRegistrada) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparando consulta SQL: %v", err)
	}
	defer query.Close() // Aquí sí se usa defer correctamente

	_, err = query.Exec(oxygen.Id_user, fecha.Format("2006-01-02"), hora.Format("15:04:05"), oxygen.RegisteredMeasure)
	if err != nil {
		log.Println("Error guardando el oxígeno medido:", err)
		return fmt.Errorf("error ejecutando la consulta SQL: %v", err)
	}

	log.Println("Oxígeno guardado correctamente en la base de datos.")
	return nil
}
func (sql *MySQL) GetUserOxygen() ([]domain.UserOxygen, error) {
	var userOxygenations []domain.UserOxygen

	query := `
    SELECT 
        ro.id_oxigeno,ro.fecha,ro.hora,ro.medidaRegistrada,
        u.id_usuario,u.nombre,u.correo,u.password,u.premium
    FROM Usuario u INNER JOIN registrooxigeno ro ON ro.id_user = u.id_usuario`

	rows, err := sql.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var uo domain.UserOxygen

		err := rows.Scan(
			&uo.Id_oxygen,
			&uo.Date,
			&uo.Time,
			&uo.RegisteredMeasure,
			&uo.Id_user,
			&uo.Name,
			&uo.Email,
			&uo.Password,
			&uo.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userOxygenations = append(userOxygenations, uo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userOxygenations, nil

}
