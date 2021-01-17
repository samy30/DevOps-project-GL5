package models

// Transaction ...
type Transaction struct {
	Date     string   `json:"date,omitempty" bson:"date,omitempty"`
	Product  *Product `json:"product" bson:"product,omitempty"`
	Quantity int      `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
