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

type MongoDBDataStore struct {
	Url            string
	DatabaseName   string
	CollectionName string
	client         *mongo.Client
	collection     *mongo.Collection
	ctx            context.Context
}

func (d *MongoDBDataStore) Init() error {
	d.ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(d.Url)
	client, err := mongo.Connect(d.ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(d.ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	d.collection = client.Database(d.DatabaseName).Collection(d.CollectionName)
	return nil
}

func (d *MongoDBDataStore) GetItem(id string) (model.Item, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return model.Item{}, err
	}

	singleResult := d.collection.FindOne(d.ctx, bson.M{"_id": ID})

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

func (d *MongoDBDataStore) GetItems() []model.Item {
	var cur, err = d.collection.Find(d.ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	var results []model.Item
	for cur.Next(d.ctx) {
		var elem model.Item
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}

		results = append(results, elem)
	}

	cur.Close(d.ctx)

	return results
}

func (d *MongoDBDataStore) AddItem(item model.Item) (model.Item, error) {
	insertResult, err := d.collection.InsertOne(d.ctx, item)
	if err != nil {
		log.Println(err)
		return item, err
	}

	id := insertResult.InsertedID
	item.Id = id.(primitive.ObjectID).Hex()
	return item, nil
}

func (d *MongoDBDataStore) UpdateItem(item model.Item, id string) error {

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": ID}

	_, err = d.collection.ReplaceOne(d.ctx, filter, item)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
