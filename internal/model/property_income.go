package model

type PropertyIncome struct {
	ID                  string `json:"name"`
	PropertyIncomeName  string `json:"property_income_name"`
	PropertyIncomePrice string `json:"property_income_name"`
	Category            string `json:"category"`
	Description         string `json:"description"`
	Properties_id       string `json:"properties_id"`
}
