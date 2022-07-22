package main

import (
	"go_databases/Mongo"
)

func main() {
	connector := Mongo.NewMongoConnection()

	connector.Ping()
	connector.InsertDocument()
}
