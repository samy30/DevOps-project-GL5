package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// omitempty means "neglect attribute when empty during the serialization process"
// bson is a binary version of json used by mongodb to improve the process of serialization and filters

// Product ...
type Product struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title           string             `json:"title" bson:"title,omitempty"`
	Price           float32            `json:"price,omitempty" bson:"price,omitempty"`
	Quantity        int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	InitialQuantity int                `json:"initial_quantity,omitempty" bson:"initial_quantity,omitempty"`
	Category        *Category          `json:"category" bson:"category,omitempty"`
}

// Category ...
type Category struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// Transaction ...
type Transaction struct {
	Date     string   `json:"date,omitempty" bson:"date,omitempty"`
	Product  *Product `json:"product" bson:"product,omitempty"`
	Quantity int      `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
