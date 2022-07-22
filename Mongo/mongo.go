package Mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Database interface {
	Ping()
	InsertDocument()
}

type MongoConnection struct{}

func (m *MongoConnection) connect() (context.Context, *mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err == nil {
		fmt.Println("Connect Success")
	} else {
		fmt.Printf("Connect Error %s\n", err)
	}
	return ctx, client, cancel
}

func disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		panic(err)
	}
}

func (m *MongoConnection) Ping() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)

	err := client.Ping(ctx, readpref.Primary())
	if err == nil {
		fmt.Println("Ping Success")
	} else {
		fmt.Printf("Ping Error %s\n", err)
	}
}

func (m *MongoConnection) InsertDocument() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)

	collection := client.Database("Music").Collection("record")
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})

	if err != nil {
		fmt.Printf("Insert Error: %s\n", err)
	} else {
		id := res.InsertedID
		fmt.Printf("Inserted Document: %s\n", id)
	}
}

func NewMongoConnection() Database {
	return &MongoConnection{}
}
