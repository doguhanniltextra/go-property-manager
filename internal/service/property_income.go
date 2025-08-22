package service

import (
	"database/sql"
	"fmt"

	"github.com/doguhanniltextra/property_go/internal/model"
)

func PropertyIncomeCreate(db *sql.DB, property_income *model.PropertyIncome) (sql.Result, error) {
	result, err := db.Exec(`
	INSERT INTO propertyincomes
	( propertyincomename, propertyincomeprice, category,description,properties_id)
	VALUES ($1, $2, $3, $4, $5)
	`, property_income.PropertyIncomeName, property_income.PropertyIncomePrice, property_income.Category, property_income.Description, property_income.Properties_id)

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	fmt.Printf("rows: %v\n", rows)

	return result, nil
}
