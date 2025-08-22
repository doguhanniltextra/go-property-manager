package service

import (
	"database/sql"
	"fmt"

	"github.com/doguhanniltextra/property_go/internal/model"
	"github.com/sirupsen/logrus"
)

func CreateProperty(db *sql.DB, property *model.Property, user_id int) (sql.Result, error) {
	result, err := db.Exec(`
	INSERT INTO properties 
	(name, purchaseprice, purchasedate, address, propertytype, areasqm, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, property.Name, property.PurchasePrice, property.PurchaseDate, property.Address, property.PropertyType, property.AreaSqm, user_id)

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	logrus.Infof("Rows affected: %d", rows)
	logrus.Infof("userid -> %d ", user_id)

	return result, nil
}

func DeleteProperty(db *sql.DB, id int) error {
	result, err := db.Exec(`
		DELETE FROM properties
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no property found with id %d", id)
	}

	return nil
}

func UpdateProperty(db *sql.DB, givenProperty *model.PropertyUpdate) error {
	_, err := db.Exec(`
		UPDATE properties
		SET Name=$1, PurchasePrice=$2, PurchaseDate=$3, Address=$4, PropertyType=$5, AreaSqm=$6
		WHERE id=$7
	`, givenProperty.Name, givenProperty.PurchasePrice, givenProperty.PurchaseDate,
		givenProperty.Address, givenProperty.PropertyType, givenProperty.AreaSqm, givenProperty.ID)

	if err != nil {
		return err
	}

	return err
}
