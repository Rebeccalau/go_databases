package main

import (
	"go_databases/mongo_implementation"
	"go_databases/sqlite_implementation"
)

func main() {
	connector := sqlite_implementation.NewSqlConnection()

	connector.Ping()
	//connector.CreateTable()
	//connector.InsertRow()
	//connector.DeleteRow()
	//connector.UpdateRow()
	connector.SearchAll()
}

func NoSQL() {
	connector := mongo_implementation.NewMongoConnection()
	//connector := redis_implementation.NewRedisConnection()

	connector.Ping()
	connector.InsertDocument()
	connector.DeleteManyUsingFilter()
	connector.DeleteDocument()
	connector.SearchAllDocument()
	connector.UpdateDocument()
}
