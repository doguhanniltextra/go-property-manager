package service

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/internal/middleware"
	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/sirupsen/logrus"
)

func RegisterService(db *sql.DB, givenUser *model.User) (string, error) {

	logrus.Info("Register Service Starting")

	var id int
	err := db.QueryRow(`
        INSERT INTO users(name,email,password) 
        VALUES($1,$2,$3) RETURNING id
    `, givenUser.Name, givenUser.Email, givenUser.Password).Scan(&id)

	givenUser.ID = id

	if err != nil {
		return "", nil
	}

	createdToken, err := middleware.CreateToken(givenUser)
	if err != nil {
		logrus.Info("Create Token didn't work.")
		return "", nil
	}

	logrus.Info("Register Service is completed.")

	return createdToken, nil
}

func AuthService(db *sql.DB, authRequest *model.AuthRequest) (bool, error) {
	row := db.QueryRow(`
        SELECT id 
        FROM users
        WHERE name = $1 AND password = $2
    `, authRequest.Name, authRequest.Password)

	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAllUsers(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(`
		SELECT id, name, email, password FROM users
	`)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
