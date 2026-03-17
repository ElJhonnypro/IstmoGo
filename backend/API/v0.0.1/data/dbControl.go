package data

// Connect to postgresSQL

import (
	carsUseModels "API/internal/carmanagement/models"
	rideModels "API/internal/ridesmanagement/models"
	userUseModels "API/internal/usermanagement/models"
	"API/internal/utils"
	"database/sql"
	"log"
	"os"

	"errors"

	_ "github.com/lib/pq"
)

// System

var ErrRideNotAvailable = errors.New("ride is not available for acceptance")

func executeQuery(db *sql.DB, query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) {

	// CARS primero
	queryCars := `
	CREATE TABLE IF NOT EXISTS public.cars (
		id TEXT PRIMARY KEY,
		plate VARCHAR(15) UNIQUE,
		model VARCHAR(50),
		color VARCHAR(30),
		photo VARCHAR(255)
	);
	`
	if _, err := db.Exec(queryCars); err != nil {
		log.Fatal("Error creating cars table:", err)
	}

	// USERS después
	queryUsers := `
	CREATE TABLE IF NOT EXISTS public.users (
		id TEXT PRIMARY KEY,
		name VARCHAR(32),
		phone VARCHAR(15) UNIQUE NOT NULL,
		role VARCHAR(50),
		rid VARCHAR(50),
		rid_photo VARCHAR(255),
		car_id TEXT NULL REFERENCES public.cars(id),
		email VARCHAR(100) UNIQUE,
		password VARCHAR(255),
		location VARCHAR(255),
		birthdate DATE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
		`
	if _, err := db.Exec(queryUsers); err != nil {
		log.Fatal("Error creating users table:", err)
	}

	queryRider := `
	CREATE TABLE IF NOT EXISTS public.rides(
		id TEXT PRIMARY KEY,
		client_id TEXT NOT NULL REFERENCES users(id),
		driver_id TEXT REFERENCES users(id),
		

		start_lat DOUBLE PRECISION NOT NULL,
		start_lng DOUBLE PRECISION NOT NULL,
		end_lat DOUBLE PRECISION NOT NULL,
		end_lng DOUBLE PRECISION NOT NULL,

		distance_km NUMERIC(10,2),
		price NUMERIC(10,2),

		status TEXT NOT NULL DEFAULT 'requested',

		created_at TIMESTAMP DEFAULT NOW(),
		requested_at TIMESTAMP DEFAULT NOW(),
		accepted_at TIMESTAMP,
		started_at TIMESTAMP,
		finished_at TIMESTAMP
		
	)
		
	`

	db.Exec(queryRider)

}

var DB *sql.DB

func ConnectDB() {
	connURL := os.Getenv("DB_URL")

	var err error
	DB, err = sql.Open("postgres", connURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}

var db = GetDB()

// Funcs

// Success need to be obligatory

type InsertUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type InsertCarResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Id      string `json:"carId"`
}

// Users

func InsertUser(db *sql.DB, user userUseModels.User) InsertUserResponse {
	// Query to insert user
	query := `
	INSERT INTO users (id,name, phone, role, rid, rid_photo, car_id, email, password, location, birthdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`
	_, err := db.Exec(query,
		user.ID,
		user.Name,
		user.Phone,
		user.Role,
		user.RID,
		user.RIDPhoto,
		user.CarID,
		user.Email,
		user.Password,
		user.Location,
		user.Birthdate)

	if err != nil {
		log.Println(err)
		return InsertUserResponse{
			Success: false,
			Message: "Failed to insert user: " + err.Error(),
		}
	}

	return InsertUserResponse{
		Success: true,
		Message: "User inserted successfully",
	}
}

func AdmingetUsers(db *sql.DB) ([]userUseModels.User, error) {
	query := `SELECT id, name, phone, role, rid, rid_photo, car_id, email, password, location, birthdate, created_at FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userUseModels.User
	for rows.Next() {
		var user userUseModels.User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Role, &user.RID, &user.RIDPhoto, &user.CarID, &user.Email, &user.Password, &user.Location, &user.Birthdate, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func NormalgetUsers(db *sql.DB) ([]userUseModels.User, error) {
	query := `SELECT id, name, phone, role, rid, car_id, email, birthdate, created_at FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []userUseModels.User
	for rows.Next() {
		var user userUseModels.User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Role, &user.RID, &user.CarID, &user.Email, &user.Birthdate, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(db *sql.DB, admin bool, userID string) (userUseModels.User, error) {
	var user userUseModels.User

	query := `SELECT`

	if admin {
		query += ` id, name, phone, role, rid, rid_photo, car_id, email, password, location, birthdate, created_at `
	} else {
		query += ` id, name, phone, role, rid, car_id, email, birthdate, created_at `
	}

	query += ` FROM users WHERE id = $1;`

	row := db.QueryRow(query, userID)

	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Role, &user.RID, &user.RIDPhoto, &user.CarID, &user.Email, &user.Password, &user.Location, &user.Birthdate, &user.CreatedAt)
	return user, err

}

func VerifyLogIn(db *sql.DB, infoValue string) (userUseModels.User, error) {
	var user userUseModels.User
	var query string

	if utils.IsEmail(infoValue) {
		query = `
		SELECT id, password
		FROM users
		WHERE email = $1;
		`
	} else if utils.IsPhoneNumber(infoValue) {
		query = `
		SELECT id, password
		FROM users
		WHERE phone = $1;
		`
	} else {
		return user, errors.New("invalid login identifier")
	}

	err := db.QueryRow(query, infoValue).Scan(
		&user.ID,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("invalid credentials")
		}
		return user, err
	}

	return user, nil
}

// Cars

func InsertCar(db *sql.DB, car carsUseModels.Car) InsertCarResponse {
	query := `
	INSERT INTO cars (id,plate, model, color, photo)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err := db.Exec(query,
		car.ID,
		car.Plate,
		car.Model,
		car.Color,
		car.Photo)

	if err != nil {
		log.Println(err)
		return InsertCarResponse{
			Success: false,
			Message: "Failed to insert car: " + err.Error(),
		}
	}

	return InsertCarResponse{
		Success: true,
		Message: "Car inserted successfully",
		Id:      car.ID,
	}
}

func GetCardById(db *sql.DB, Id string) (carsUseModels.Car, error) {
	var car carsUseModels.Car

	query := "SELECT * FROM cars WHERE id=$1"
	row := db.QueryRow(query, Id)

	err := row.Scan(&car.ID, &car.Plate, &car.Model, &car.Color, &car.Photo)
	return car, err

}

func InsertRide(db *sql.DB, ride rideModels.Ride) error {
	query := `
	INSERT INTO rides (
		id, client_id,
		start_lat, start_lng,
		end_lat, end_lng,
		distance_km, price, status
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,'requested')
	`

	_, err := db.Exec(
		query,
		ride.ID,
		ride.ClientID,
		ride.StartLat,
		ride.StartLng,
		ride.EndLat,
		ride.EndLng,
		ride.DistanceKm,
		ride.Price,
	)

	return err
}

func GetAllRides(db *sql.DB) ([]rideModels.Ride, error) {
	query := `SELECT id, client_id, driver_id, car_id, start_lat, start_lng, end_lat, end_lng, distance_km, price, status, requested_at, accepted_at, started_at, finished_at FROM rides;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rides []rideModels.Ride
	for rows.Next() {
		var ride rideModels.Ride
		err := rows.Scan(
			&ride.ID,
			&ride.ClientID,
			&ride.UberID,
			&ride.StartLat,
			&ride.StartLng,
			&ride.EndLat,
			&ride.EndLng,
			&ride.DistanceKm,
			&ride.Price,
			&ride.Status,
			&ride.RequestedAt,
			&ride.AcceptedAt,
			&ride.StartedAt,
			&ride.FinishedAt,
		)
		if err != nil {
			return nil, err
		}
		rides = append(rides, ride)
	}
	return rides, nil
}

func GetRideByID(db *sql.DB, rideID string) (rideModels.Ride, error) {
	var ride rideModels.Ride
	query := `SELECT id, client_id, driver_id, car_id, start_lat, start_lng, end_lat, end_lng, distance_km, price, status, requested_at, accepted_at, started_at, finished_at FROM rides WHERE id = $1;`
	row := db.QueryRow(query, rideID)
	err := row.Scan(
		&ride.ID,
		&ride.ClientID,
		&ride.UberID,
		&ride.StartLat,
		&ride.StartLng,
		&ride.EndLat,
		&ride.EndLng,
		&ride.DistanceKm,
		&ride.Price,
		&ride.Status,
		&ride.RequestedAt,
		&ride.AcceptedAt,
		&ride.StartedAt,
		&ride.FinishedAt,
	)
	return ride, err
}

func UpdateRide(db *sql.DB, ride rideModels.Ride) error {
	query := `
	UPDATE rides
	SET client_id = $1,
		driver_id = $2,
		start_lat = $3,
		start_lng = $4,
		end_lat = $5,
		end_lng = $6,
		distance_km = $7,
		price = $8,
		status = $9,
		requested_at = $10,
		accepted_at = $11,
		started_at = $12,
		finished_at = $13
	WHERE id = $14;
	`
	_, err := db.Exec(
		query,
		ride.ClientID,
		ride.UberID,
		ride.StartLat,
		ride.StartLng,
		ride.EndLat,
		ride.EndLng,
		ride.DistanceKm,
		ride.Price,
		ride.Status,
		ride.RequestedAt,
		ride.AcceptedAt,
		ride.StartedAt,
		ride.FinishedAt,
		ride.ID,
	)
	return err
}

func GetActiveRidesByUserID(db *sql.DB, userID string) ([]rideModels.Ride, error) {
	query := `SELECT id, client_id, driver_id, car_id, start_lat, start_lng, end_lat, end_lng, distance_km, price, status, requested_at, accepted_at, started_at, finished_at FROM rides WHERE (client_id = $1 OR driver_id = $1) AND status IN ('requested', 'accepted', 'started');`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rides []rideModels.Ride
	for rows.Next() {
		var ride rideModels.Ride
		err := rows.Scan(
			&ride.ID,
			&ride.ClientID,
			&ride.UberID,
			&ride.StartLat,
			&ride.StartLng,
			&ride.EndLat,
			&ride.EndLng,
			&ride.DistanceKm,
			&ride.Price,
			&ride.Status,
			&ride.RequestedAt,
			&ride.AcceptedAt,
			&ride.StartedAt,
			&ride.FinishedAt,
		)
		if err != nil {
			return nil, err
		}
		rides = append(rides, ride)
	}
	return rides, nil
}
