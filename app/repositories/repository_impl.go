package repositories

import (
	"context"

	"devopsProjectModule.com/gl5/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	db *mongo.Collection
}

// constructing the repositiory
func NewProductRepository(db *mongo.Collection) Repository {
	return &productRepository{
		db: db,
	}
}

// Get all the products from the database.
func (r productRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	cur, err := r.db.Find(ctx, bson.M{})

	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		product := new(models.Product)
		err := cur.Decode(product)

		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

// Get product with the specified ID from the database.
func (r productRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product *models.Product
	// string to primitive.ObjectID
	productId, _ := primitive.ObjectIDFromHex(id)

	err := r.db.FindOne(ctx, bson.M{"_id": productId}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, err
}

// Create product with the specified object
func (r productRepository) Create(ctx context.Context, product *models.Product) (*primitive.ObjectID, error) {
	result, err := r.db.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	insertedID, _ := result.InsertedID.(primitive.ObjectID)
	return &insertedID, nil

}

// Update saves the changes to a product in the database.
func (r productRepository) Update(ctx context.Context, product *models.Product) error {
	id, _ := primitive.ObjectIDFromHex(product.ID)

	// prepare update model.
	fieldsToUpdate := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "title", Value: product.Title},
				primitive.E{Key: "price", Value: product.Price},
				primitive.E{Key: "initial_quantity", Value: product.InitialQuantity},
				primitive.E{Key: "quantity", Value: product.Quantity},
				primitive.E{Key: "category", Value: bson.D{
					primitive.E{Key: "name", Value: product.Category.Name},
				}},
			},
		},
	}

	_, err := r.db.UpdateOne(ctx, bson.M{"_id": id}, fieldsToUpdate)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a product with the specified ID from the database.
func (r productRepository) Delete(ctx context.Context, id string) error {
	productId, _ := primitive.ObjectIDFromHex(id)
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": productId})
	if err != nil {
		return err
	}
	return nil
}
