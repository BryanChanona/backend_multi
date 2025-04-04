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

	query, err := sql.db.Prepare("INSERT INTO registrooxigeno (id_user, fecha, hora, medidaRegistrada,id_dispositivo) VALUES (?, ?, ?, ?,?)")
	if err != nil {
		return fmt.Errorf("error preparando consulta SQL: %v", err)
	}
	defer query.Close() // Aquí sí se usa defer correctamente

	_, err = query.Exec(oxygen.Id_user, fecha.Format("2006-01-02"), hora.Format("15:04:05"), oxygen.RegisteredMeasure,oxygen.Id_device)
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
        u.id_usuario,u.nombre,u.correo,u.premium
    FROM usuario u INNER JOIN registrooxigeno ro ON ro.id_user = u.id_usuario`

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
func (sql *MySQL)GetOxygenByDate( date string,idUser int)([]domain.UserOxygen, error){
	var userOxygenArray []domain.UserOxygen

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha: %v", err)
	}

	query := `SELECT 
        ro.id_oxigeno,ro.fecha,ro.hora,ro.medidaRegistrada,
        u.id_usuario,u.nombre,u.correo,u.premium
    FROM usuario u INNER JOIN registrooxigeno ro ON ro.id_user = u.id_usuario WHERE ro.fecha =? AND u.id_usuario = ?`
	rows, err := sql.db.Query(query, parsedDate, idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var userOxygen domain.UserOxygen

		// Escanear los resultados de cada fila
		err := rows.Scan(
			&userOxygen.Id_oxygen,
			&userOxygen.Date,
			&userOxygen.Time,
			&userOxygen.RegisteredMeasure,
			&userOxygen.Id_user,
			&userOxygen.Name,
			&userOxygen.Email,
			&userOxygen.Premium)

		// Si ocurre un error al escanear la fila
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		// Añadir el registro al array
		userOxygenArray = append(userOxygenArray, userOxygen)

	}

	// Revisar si hubo algún error durante el recorrido de las filas
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}



	// Retornar el array de resultados
	return userOxygenArray, nil

	

	

}
func (sql *MySQL)GetOxygenById(idUser int) ([]domain.UserOxygen, error){
	var userOxygens []domain.UserOxygen

	query := `
	SELECT 
        ro.id_oxigeno,ro.fecha,ro.hora,ro.medidaRegistrada,
        u.id_usuario,u.nombre,u.correo,u.premium
    FROM usuario u INNER JOIN registrooxigeno ro ON ro.id_user = u.id_usuario WHERE u.id_usuario = ?`

	rows, err := sql.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var userOxygen domain.UserOxygen

		err := rows.Scan(
			&userOxygen.Id_oxygen,
			&userOxygen.Date,
			&userOxygen.Time,
			&userOxygen.RegisteredMeasure,
			&userOxygen.Id_user,
			&userOxygen.Name,
			&userOxygen.Email,
			&userOxygen.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userOxygens = append(userOxygens, userOxygen)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userOxygens, nil
}

func (sq *MySQL)GetOxygenSupervisorByIdUser(idUser int)([]domain.UserOxygen, error){
	var userOxygenArray []domain.UserOxygen

	query := `SELECT 
		ro.id_oxigeno,ro.fecha,ro.hora,ro.medidaRegistrada,
		u.id_usuario,u.nombre,u.correo,u.premium
	FROM usuario u INNER JOIN registrooxigeno ro ON ro.id_user = u.id_usuario WHERE u.id_usuario = ?`

	rows, err := sq.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userOxygenSupervisor domain.UserOxygen

		err := rows.Scan(
			&userOxygenSupervisor.Id_oxygen,
			&userOxygenSupervisor.Date,
			&userOxygenSupervisor.Time,
			&userOxygenSupervisor.RegisteredMeasure,
			&userOxygenSupervisor.Id_user,
			&userOxygenSupervisor.Name,
			&userOxygenSupervisor.Email,
			&userOxygenSupervisor.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userOxygenArray = append(userOxygenArray, userOxygenSupervisor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userOxygenArray, nil
}