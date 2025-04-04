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

	if temperatura.Date == "" {
        temperatura.Date = time.Now().Format("2006-01-02")  // Formato YYYY-MM-DD
    }

    // Si no se proporciona hora, usa la hora actual
    if temperatura.Time == "" {
        temperatura.Time = time.Now().Format("15:04")  // Formato HH:mm
    }

    // Parsear la fecha
    fecha, err := time.Parse("2006-01-02", temperatura.Date)
    if err != nil {
        return fmt.Errorf("formato de fecha inválido: %v", err)
    }

    // Parsear la hora
    hora, err := time.Parse("15:04", temperatura.Time)
    if err != nil {
        return fmt.Errorf("formato de hora inválido: %v", err)
    }

    // Preparar la consulta SQL para insertar
    query := `INSERT INTO registrotemperatura 
              (id_user, fecha, hora, medidaregistrada,id_dispositivo) 
              VALUES (?, ?, ?, ?,?)`
    
    _, err = sql.db.Exec(query,
        temperatura.Id_user,
        fecha.Format("2006-01-02"),
        hora.Format("15:04:00"),
        temperatura.RegisteredMeasure,temperatura.Id_device)
    
    if err != nil {
        return fmt.Errorf("error al guardar temperatura: %v", err)
    }
    
    return nil
}
func (sql *MySQL) GetTemperature() ([]domain.UserTemperature, error) {
	var userTemperatures []domain.UserTemperature

	query := `
		SELECT 
			rt.id_temp, u.id_usuario, u.nombre, u.correo, u.premium, 
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
			&ut.Email,              // STRING (correo)          // STRING (password)
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

func (sql *MySQL) GetTemperatureByDate( idUser int,date string,) ([]domain.UserTemperature, error) {
	var userTemperatureArray []domain.UserTemperature

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha: %v", err)
	}

	query := `
	SELECT 
			rt.id_temp, u.id_usuario, u.nombre, u.correo, u.premium, 
			rt.medidaRegistrada, rt.fecha, rt.hora 
		FROM Usuario u 
		INNER JOIN RegistroTemperatura rt ON rt.id_user = u.id_usuario 
		WHERE rt.fecha = ? AND u.id_usuario = ?`

	rows, err := sql.db.Query(query, parsedDate, idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var userTemperature domain.UserTemperature

		// Escanear los resultados de cada fila
		err := rows.Scan(
			&userTemperature.Id_temp,
		&userTemperature.Id_user,
		&userTemperature.Name,
		&userTemperature.Email,
		&userTemperature.Premium,
		&userTemperature.RegisteredMeasure,
		&userTemperature.Date,
		&userTemperature.Time)

		// Si ocurre un error al escanear la fila
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		// Añadir el registro al array
		userTemperatureArray= append(userTemperatureArray, userTemperature)

	}

	// Revisar si hubo algún error durante el recorrido de las filas
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}



	// Retornar el array de resultados
	return userTemperatureArray, nil
}

func (sql *MySQL)GetTemperatureById(idUser int)([]domain.UserTemperature, error){
	var userTemperatures []domain.UserTemperature

	query:= `
	SELECT 
			rt.id_temp, u.id_usuario, u.nombre, u.correo, u.premium, 
			rt.medidaRegistrada, rt.fecha, rt.hora 
		FROM Usuario u 
		INNER JOIN RegistroTemperatura rt ON rt.id_user = u.id_usuario 
		WHERE u.id_usuario = ?`
	
		rows, err := sql.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		var userTemperature domain.UserTemperature
		 err := rows.Scan(
			&userTemperature.Id_temp,
		&userTemperature.Id_user,
		&userTemperature.Name,
		&userTemperature.Email,
		&userTemperature.Premium,
		&userTemperature.RegisteredMeasure,
		&userTemperature.Date,
		&userTemperature.Time)
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}
		userTemperatures = append(userTemperatures,userTemperature)


	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userTemperatures, nil
	

}

func (sql *MySQL)GetTemperatureSupervisorByIdUser(idUser int) ([]domain.UserTemperature,error){
	var userTemperatureArray []domain.UserTemperature

	query := `SELECT 
		rt.id_temp,rt.fecha,rt.hora,rt.medidaRegistrada,
		u.id_usuario,u.nombre, u.correo,u.premium
	FROM Usuario u INNER JOIN registrotemperatura rt ON rt.id_user = u.id_usuario WHERE  u.id_usuario = ?`

	rows, err := sql.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userTemperatureSupervisor domain.UserTemperature

		err := rows.Scan(
			&userTemperatureSupervisor.Id_temp,
			&userTemperatureSupervisor.Date,
			&userTemperatureSupervisor.Time,
			&userTemperatureSupervisor.RegisteredMeasure,
			&userTemperatureSupervisor.Id_user,
			&userTemperatureSupervisor.Name,
			&userTemperatureSupervisor.Email,
			&userTemperatureSupervisor.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userTemperatureArray = append(userTemperatureArray, userTemperatureSupervisor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userTemperatureArray, nil
}