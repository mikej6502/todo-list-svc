package database

import (
	"context"
	"errors"
	"github.com/mikej6502/todo-list-svc/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client
var collection *mongo.Collection
var ctx = context.TODO()

type MongoDBDataStore struct {
	Url string
}

func (d MongoDBDataStore) Init() error {
	clientOptions := options.Client().ApplyURI(d.Url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	collection = client.Database("todo-list-db").Collection("items")
	return nil
}

func (d MongoDBDataStore) GetItem(id string) (model.Item, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return model.Item{}, err
	}

	filter := bson.M{"_id": ID}

	findOptions := options.Find()
	findOptions.SetLimit(1)

	singleResult := collection.FindOne(ctx, filter)

	if singleResult == nil {
		log.Println("item not found for ID: " + id)
		return model.Item{}, errors.New("item not found for ID: " + id)
	}

	var elem model.Item
	err = singleResult.Decode(&elem)
	if err != nil {
		log.Println(err)
		return model.Item{}, err
	}

	return elem, nil
}

func (d MongoDBDataStore) GetItems() []model.Item {
	var cur, err = collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	var results []model.Item
	for cur.Next(ctx) {
		var elem model.Item
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}

		results = append(results, elem)
	}

	cur.Close(ctx)

	return results
}

func (d MongoDBDataStore) AddItem(item model.Item) (model.Item, error) {
	insertResult, err := collection.InsertOne(ctx, item)
	if err != nil {
		log.Println(err)
		return item, err
	}

	id := insertResult.InsertedID
	item.Id = id.(primitive.ObjectID).Hex()
	return item, nil
}

func (d MongoDBDataStore) UpdateItem(item model.Item, id string) error {

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": ID}

	_, err = collection.ReplaceOne(ctx, filter, item)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
