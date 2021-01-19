package models

type Product struct {
	ID              string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title           string    `json:"title" bson:"title,omitempty"`
	Price           float32   `json:"price,omitempty" bson:"price,omitempty"`
	Quantity        int       `json:"quantity" bson:"quantity"`
	InitialQuantity int       `json:"initial_quantity" bson:"initial_quantity"`
	Category        *Category `json:"category" bson:"category,omitempty"`
}
