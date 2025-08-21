package main

import (
	"database/sql"

	"github.com/doguhanniltextra/property_go/database"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"os"
)

type DatabaseC struct {
	DB *sql.DB
}

func main() {
	loggerInit()
	databaseVariables()

	db, err := database.Connection()
	if err != nil {
		logrus.Infoln("Database connection failed: ", err)
	}

	defer db.Close()
}

func loggerInit() {
	log.SetLevel(log.InfoLevel)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
		PadLevelText:     true,
		ForceQuote:       true,
	})

}

func databaseVariables() {
	logrus.Infoln("Environment variables loaded from .env:")
	logrus.Infoln("DB_HOST:", os.Getenv("DB_HOST"))
	logrus.Infoln("DB_USER:", os.Getenv("DB_USER"))
	logrus.Infoln("DB_NAME:", os.Getenv("DB_NAME"))
	logrus.Infoln("DB_PORT:", os.Getenv("DB_PORT"))

	logrus.Infoln("DB_PASSWORD: [HIDDEN]")
}
