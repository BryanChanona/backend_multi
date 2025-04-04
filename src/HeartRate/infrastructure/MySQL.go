package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/BryanChanona/backend_multi/src/HeartRate/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}


func (sql *MySQL) SaveHeartRate(heartRate domain.HeartRate) error{
	if heartRate.Date == "" {
        heartRate.Date = time.Now().Format("2006-01-02")  // Formato YYYY-MM-DD
    }

    // Si no se proporciona hora, usa la hora actual
    if heartRate.Time == "" {
        heartRate.Time = time.Now().Format("15:04")  // Formato HH:mm
    }

    // Parsear la fecha
    fecha, err := time.Parse("2006-01-02", heartRate.Date)
    if err != nil {
        return fmt.Errorf("formato de fecha inválido: %v", err)
    }

    // Parsear la hora
    hora, err := time.Parse("15:04", heartRate.Time)
    if err != nil {
        return fmt.Errorf("formato de hora inválido: %v", err)
    }
	query, err := sql.db.Prepare("INSERT INTO registrobpm (id_user, fecha, hora, medidaRegistrada,id_dispositivo) VALUES (?, ?, ?, ?,?)")
	if err != nil {
		return fmt.Errorf("error preparando consulta SQL: %v", err)
	}
	defer query.Close() // Aquí sí se usa defer correctamente
	_, err = query.Exec(heartRate.Id_user, fecha.Format("2006-01-02"), hora.Format("15:04:05"), heartRate.RegisteredMeasure,heartRate.Id_device)
	if err != nil {
		log.Println("Error guardando el oxígeno medido:", err)
		return fmt.Errorf("error ejecutando la consulta SQL: %v", err)
	}

	return nil

}
func (sql *MySQL)GetUserHeartRate() ([]domain.UserHeartRate, error){
	var userHeartRates []domain.UserHeartRate

	query := `
	SELECT
		rb.id_bpm,rb.fecha,rb.hora,rb.medidaRegistrada,
		u.id_usuario,u.nombre, u.correo,u.premium
	FROM Usuario u INNER JOIN registrobpm rb ON rb.id_user = u.id_usuario`

	rows, err := sql.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userHeartRate domain.UserHeartRate

		err := rows.Scan(
			&userHeartRate.Id_bpm,
			&userHeartRate.Date,
			&userHeartRate.Time,
			&userHeartRate.RegisteredMeasure,
			&userHeartRate.Id_user,
			&userHeartRate.Name,
			&userHeartRate.Email,
			&userHeartRate.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userHeartRates = append(userHeartRates, userHeartRate)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userHeartRates, nil



}

func (sql *MySQL)GetHeartRateByDate(idUser int,date string) ([]domain.UserHeartRate,error){
	var userHeartRateArray []domain.UserHeartRate

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha: %v", err)
	}

	query := `SELECT 
       rb.id_bpm,rb.fecha,rb.hora,rb.medidaRegistrada,
		u.id_usuario,u.nombre, u.correo,u.premium
	FROM Usuario u INNER JOIN registrobpm rb ON rb.id_user = u.id_usuario WHERE rb.fecha =? AND u.id_usuario = ?`

	rows, err := sql.db.Query(query, parsedDate, idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var userHeartRate domain.UserHeartRate

		// Escanear los resultados de cada fila
		err := rows.Scan(
			&userHeartRate.Id_bpm,
			&userHeartRate.Date,
			&userHeartRate.Time,
			&userHeartRate.RegisteredMeasure,
			&userHeartRate.Id_user,
			&userHeartRate.Name,
			&userHeartRate.Email,
			&userHeartRate.Premium)

		// Si ocurre un error al escanear la fila
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		// Añadir el registro al array
		userHeartRateArray = append(userHeartRateArray, userHeartRate)

	}

	// Revisar si hubo algún error durante el recorrido de las filas
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}



	// Retornar el array de resultados
	return userHeartRateArray, nil
}
func (sql *MySQL) GetHeartRateById(idUser int) ([]domain.UserHeartRate,error){
	var userHeartRatesTwo []domain.UserHeartRate

	query := `SELECT 
       rb.id_bpm,rb.fecha,rb.hora,rb.medidaRegistrada,
		u.id_usuario,u.nombre, u.correo,u.premium
	FROM Usuario u INNER JOIN registrobpm rb ON rb.id_user = u.id_usuario WHERE  u.id_usuario = ?`

	rows, err := sql.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userHeartRate domain.UserHeartRate

		err := rows.Scan(
			&userHeartRate.Id_bpm,
			&userHeartRate.Date,
			&userHeartRate.Time,
			&userHeartRate.RegisteredMeasure,
			&userHeartRate.Id_user,
			&userHeartRate.Name,
			&userHeartRate.Email,
			&userHeartRate.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userHeartRatesTwo = append(userHeartRatesTwo, userHeartRate)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userHeartRatesTwo, nil


}

func (sql *MySQL)GetHeartRateSupervisorByIdUser(idUser int) ([]domain.UserHeartRate,error){
	var userHeartRatesArray []domain.UserHeartRate

	query := `SELECT 
       rb.id_bpm,rb.fecha,rb.hora,rb.medidaRegistrada,
		u.id_usuario,u.nombre, u.correo,u.premium
	FROM Usuario u INNER JOIN registrobpm rb ON rb.id_user = u.id_usuario WHERE  u.id_usuario = ?`

	rows, err := sql.db.Query(query,idUser)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userHeartRateSupervisor domain.UserHeartRate

		err := rows.Scan(
			&userHeartRateSupervisor.Id_bpm,
			&userHeartRateSupervisor.Date,
			&userHeartRateSupervisor.Time,
			&userHeartRateSupervisor.RegisteredMeasure,
			&userHeartRateSupervisor.Id_user,
			&userHeartRateSupervisor.Name,
			&userHeartRateSupervisor.Email,
			&userHeartRateSupervisor.Premium)

		if err != nil {
			return nil, fmt.Errorf("error al escanear fila: %v", err)
		}

		userHeartRatesArray = append(userHeartRatesArray,userHeartRateSupervisor)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer filas: %v", err)
	}

	return userHeartRatesArray, nil


}
