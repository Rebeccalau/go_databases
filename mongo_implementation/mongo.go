package mongo_implementation

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go_databases/records"
	"log"
	"time"
)

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

func (m *MongoConnection) collection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("Records").Collection("documents")
	return collection
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

	//doc, err := bson.Marshal(records.DefaultRecord2())
	doc, err := bson.Marshal(records.DefaultRecord1())

	if err != nil {
		fmt.Printf("Error Marshalling Doc: %s\n", err)
	}

	collection := m.collection(client)
	res, err := collection.InsertOne(ctx, doc)

	if err != nil {
		fmt.Printf("Insert Error: %s\n", err)
	} else {
		id := res.InsertedID
		fmt.Printf("Inserted Document: %s\n", id)
	}
}

func (m *MongoConnection) SearchAllDocument() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)
	collection := m.collection(client)

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		elem := &records.Record{}
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *MongoConnection) DeleteManyUsingFilter() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)

	collection := m.collection(client)

	result, err := collection.DeleteMany(ctx, bson.M{"name": "pi"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully deleted %v\n", result.DeletedCount)
}

func (m *MongoConnection) UpdateDocument() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)
	collection := m.collection(client)
	filter := bson.D{{"id", "1234"}}
	update := bson.D{{"$set", bson.D{{"secondary", "22222"}}}}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully updated %s\n", result)
}

func (m *MongoConnection) DeleteDocument() {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer disconnect(client, ctx)

	collection := m.collection(client)

	result, err := collection.DeleteOne(ctx, bson.M{"value": 3.14159})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully deleted %v\n", result.DeletedCount)
}

func NewMongoConnection() records.NoSQLDatabase {
	return &MongoConnection{}
}
