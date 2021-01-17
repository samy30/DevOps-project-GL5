package payload

type BuyRequest struct {
	ProductId string `json:"product_id" bson:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
