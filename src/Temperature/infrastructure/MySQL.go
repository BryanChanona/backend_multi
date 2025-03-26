package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/BryanChanona/backend_multi/src/Temperature/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) SaveTemperature(temperatura domain.Temperature) error {

	fecha, err := time.Parse("02/01/2006", temperatura.Date)
	if err != nil {
		return fmt.Errorf("formato de fecha inválido: %v", err)
	}

	hora, err := time.Parse("15:04", temperatura.Time)
	if err != nil {
		return fmt.Errorf("formato de hora inválido: %v", err)
	}

    // Preparar la consulta SQL para insertar
    query := `INSERT INTO registrotemperatura 
              (id_user, fecha, hora, medidaregistrada) 
              VALUES (?, ?, ?, ?)`
    
    _, err = sql.db.Exec(query,
        temperatura.Id_user,
        fecha.Format("2006-01-02"),
        hora.Format("15:04:00"),
        temperatura.RegisteredMeasure,
    )
    
    if err != nil {
        return fmt.Errorf("error al guardar temperatura: %v", err)
    }
    
    return nil
}
func (sql *MySQL) GetTemperature() ([]domain.UserTemperature, error) {
	var userTemperatures []domain.UserTemperature

	query := `
		SELECT 
			rt.id_temp, u.id_usuario, u.nombre, u.correo, u.password, u.premium, 
			rt.medidaRegistrada, rt.fecha, rt.hora 
		FROM Usuario u 
		INNER JOIN RegistroTemperatura rt ON rt.id_user = u.id_usuario
	`

	rows, err := sql.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ut domain.UserTemperature

		err := rows.Scan(
			&ut.Id_temp,            // INT
			&ut.Id_user,            // INT
			&ut.Name,               // STRING (nombre)
			&ut.Email,              // STRING (correo)
			&ut.Password,           // STRING (password)
			&ut.Premium,            // BOOL (premium)
			&ut.RegisteredMeasure,  // FLOAT (medidaRegistrada)
			&ut.Date,               // STRING (fecha)
			&ut.Time,               // STRING (hora)
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		// Agregar a la lista
		userTemperatures = append(userTemperatures, ut)
	}

	// Verificar errores al recorrer las filas
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userTemperatures, nil
}
