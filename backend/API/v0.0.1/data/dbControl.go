package data

// Connect to postgresSQL

import (
	carsUseModels "API/internal/carManagement/models"
	userUseModels "API/internal/userManagement/models"
	"API/internal/utils"
	"database/sql"
	"log"
	"os"

	"errors"

	_ "github.com/lib/pq"
)

// System

func executeQuery(db *sql.DB, query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) {
	/*
		- User Table
		id
		name
		phone
		role
		id
		rid
		rid_photo
		car_id
		email
		password
		location
		birthdate
		created_at
	*/

	/*
		- Car Table
		id
		plate
		model
		color
		photo
	*/

	// Creating User Table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name VARCHAR(32),
		phone UNIQUE NOT NULL VARCHAR(15),
		role VARCHAR(50),
		rid VARCHAR(50),
		rid_photo VARCHAR(255),
		car_id TEXT NULL REFERENCES cars(id),
		email VARCHAR(100) UNIQUE,
		password VARCHAR(255),
		location VARCHAR(255),
		birthdate DATE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);

	`

	executeQuery(db, query)

	// Creating Car Table
	query = `
	CREATE TABLE IF NOT EXISTS cars (
		id TEXT PRIMARY KEY,
		user_id TEXT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		plate VARCHAR(15) UNIQUE,
		model VARCHAR(50),
		color VARCHAR(30),
		photo VARCHAR(255)
	);
	`
	executeQuery(db, query)

}

func GetDB() *sql.DB {
	connURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Println(err)
	}
	return db
}

var db = GetDB()

// Funcs

// Success need to be obligatory

type InsertUserResponse struct {
	User    userUseModels.User `json:"user"`
	Success bool               `json:"success"`
	Message string             `json:"message"`
}

type InsertCarResponse struct {
	Car     carsUseModels.Car `json:"car"`
	Success bool              `json:"success"`
	Message string            `json:"message"`
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
		User:    user,
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

	// Decide qué campo usar
	if utils.IsEmail(infoValue) {
		query = `
		SELECT id, name, email, password, role
		FROM users
		WHERE email = $1;
		`
	} else if utils.IsPhoneNumber(infoValue) {
		query = `
		SELECT id, name, phone, password, role
		FROM users
		WHERE phone = $1;
		`
	} else if utils.IsUserName(infoValue) {
		query = `
		SELECT id, name, username, password, role
		FROM users
		WHERE username = $1;
		`
	} else {
		return user, errors.New("invalid login identifier")
	}

	// Ejecutar query
	err := db.QueryRow(query, infoValue).Scan(
		&user.ID,
		&user.Name,
		&infoValue, // email / phone / username
		&user.Password,
		&user.Role,
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
		Car:     car,
		Success: true,
		Message: "Car inserted successfully",
	}
}
