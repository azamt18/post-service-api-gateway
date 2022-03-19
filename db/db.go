package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	posts = "posts"
)

type Database interface {
	Disconnect() error
	PostsCollection() *mongo.Collection
}

type mongoDatabase struct {
	client   *mongo.Client
	database *mongo.Database

	_collections map[string]*mongo.Collection
}

func NewDatabase(connectionString string, databaseName string) Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}

	database := client.Database(databaseName)

	return &mongoDatabase{
		database:     database,
		client:       client,
		_collections: make(map[string]*mongo.Collection, 0),
	}
}

func (db *mongoDatabase) Disconnect() error {
	return db.client.Disconnect(context.TODO())
}

func (db *mongoDatabase) PostsCollection() *mongo.Collection {
	return db.collection(posts)
}

func (db *mongoDatabase) collection(name string) *mongo.Collection {
	// check for existing
	if k, v := db._collections[name]; v {
		return k
	}

	// search from database
	collectionNames, _ := db.database.ListCollectionNames(context.TODO(), bson.M{})
	for _, colName := range collectionNames {
		if colName == name {
			col := db.database.Collection(name)
			db._collections[name] = col
			return col
		}
	}

	// create collection
	db.database.CreateCollection(context.TODO(), name)
	col := db.database.Collection(name)
	db._collections[name] = col
	return col
}
