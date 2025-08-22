package model

type Property struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	PurchasePrice string `json:"purchase_price"`
	PurchaseDate  string `json:"purchase_date"`
	Address       string `json:"address"`
	PropertyType  string `json:"property_type"`
	AreaSqm       string `json:"area_sqm"`
	User_Id       int    `json:"user_id"`
}
