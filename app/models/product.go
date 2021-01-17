package models

type Product struct {
	ID              string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title           string    `json:"title" bson:"title,omitempty"`
	Price           float32   `json:"price,omitempty" bson:"price,omitempty"`
	Quantity        int       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	InitialQuantity int       `json:"initial_quantity,omitempty" bson:"initial_quantity,omitempty"`
	Category        *Category `json:"category" bson:"category,omitempty"`
}
