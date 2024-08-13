package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"product-hexagonal-architecture-go/internal/domain/entities"
	"time"
)

type ProductService struct {
	collection *mongo.Collection
}

func NewProductService(db *mongo.Database) *ProductService {
	//collections database mongo
	return &ProductService{
		collection: db.Collection("products"),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *entities.Product) error {
	product.ID = primitive.NewObjectID()
	start := time.Now()
	_, err := s.collection.InsertOne(ctx, product)
	duration := time.Since(start)
	log.Printf("CreateProduct took %s", duration)
	return err
}

func (s *ProductService) ListProducts(ctx context.Context) ([]entities.Product, error) {
	start := time.Now()
	cursor, err := s.collection.Find(ctx, bson.M{})
	duration := time.Since(start)
	log.Printf("ListProducts took %s", duration)
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	err = cursor.All(ctx, &products)
	return products, err
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	start := time.Now()
	var product entities.Product
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	duration := time.Since(start)
	log.Printf("GetProductByID took %s", duration)
	return &product, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, product *entities.Product) error {
	start := time.Now()
	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	duration := time.Since(start)
	log.Printf("UpdateProduct took %s", duration)
	return err
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	start := time.Now()
	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	duration := time.Since(start)
	log.Printf("DeleteProduct took %s", duration)
	return err
}
