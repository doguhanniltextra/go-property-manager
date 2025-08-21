package service

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/sirupsen/logrus"
)

func RegisterService(db *sql.DB, givenUser *model.User) (sql.Result, error) {

	logrus.Info("Register Service Starting")

	result, err := db.Exec(`
        INSERT INTO users (name, email, password)
         VALUES ($1, $2, $3)
    `, givenUser.Name, givenUser.Email, givenUser.Password)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logrus.Error("Failed to get rows affected: ", err)
	} else {
		logrus.Infof("Inserted user, %d row(s) affected", rowsAffected)
	}

	if err != nil {
		return nil, err
	}

	logrus.Info("Register Service is completed.")

	return result, nil
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
