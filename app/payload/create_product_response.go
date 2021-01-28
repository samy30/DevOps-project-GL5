package payload

type CreateProductResponse struct {
	ProductId string `json:"product_id" bson:"product_id,omitempty"`
}
