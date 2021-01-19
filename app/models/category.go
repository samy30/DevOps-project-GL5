package models

type Category struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
