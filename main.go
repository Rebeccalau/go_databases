package main

import (
	"go_databases/mongo_implementation"
)

func main() {
	connector := mongo_implementation.NewMongoConnection()

	connector.Ping()
	connector.InsertDocument()
	//connector.DeleteManyUsingFilter()
	//connector.DeleteDocument()
	//connector.SearchAllDocument()
	//connector.UpdateDocument()
}
