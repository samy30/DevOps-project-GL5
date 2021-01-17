package repositories

import (
	"context"

	"devopsProjectModule.com/gl5/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type transactionRepository struct {
	db *mongo.Collection
}

// constructing the repositiory
func NewTransactionRepository(db *mongo.Collection) TransactRepository {
	return &transactionRepository{
		db: db,
	}
}

// Get all the transactions from the database.
func (r transactionRepository) GetAll(ctx context.Context) ([]models.Transaction, error) {
	var transactions []models.Transaction
	cur, err := r.db.Find(ctx, bson.M{})

	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		transaction := new(models.Transaction)
		err := cur.Decode(transaction)

		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *transaction)
	}
	return transactions, err
}

// Get transaction with the specified ID from the database.
func (r transactionRepository) GetByID(ctx context.Context, id string) (models.Transaction, error) {
	var transaction models.Transaction
	// string to primitive.ObjectID
	transactionId, _ := primitive.ObjectIDFromHex(id)

	err := r.db.FindOne(ctx, bson.M{"_id": transactionId}).Decode(&transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, err
}

// Create transaction with the specified object
func (r transactionRepository) Create(ctx context.Context, transaction models.Transaction) error {
	_, err := r.db.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}
