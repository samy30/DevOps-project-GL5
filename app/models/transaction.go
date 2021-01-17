package models

// Transaction ...
type Transaction struct {
	ID       string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Date     string   `json:"date,omitempty" bson:"date,omitempty"`
	Product  *Product `json:"product" bson:"product,omitempty"`
	Quantity int      `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
