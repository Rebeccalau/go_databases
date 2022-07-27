# Mongo

## Pre Req
[Install Mongo Server](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/)
```
 brew install mongodb-community@4.4
```
*Some limitations on version for M1 Macs could not install version 6*


[Go Driver](https://github.com/mongodb/mongo-go-driver)

## Hosting 

### Run the Server
M1 Mac command for starting Mongo Server in the background
```
 mongod --config /opt/homebrew/etc/mongod.conf --fork
```
Check if mongo server is running
```
ps aux | grep -v grep | grep mongod
```
To shut down the service
- Enter the Mongo shell `mongo` 
- `use admin`
- `db.shutdownServer()`

### Run the Application
```
go run main.go
```

## FAQ
Document Versioning [Link](https://www.mongodb.com/blog/post/building-with-patterns-the-document-versioning-pattern)

An operation on a single document is atomic [Link](https://www.mongodb.com/docs/upcoming/core/transactions/)

Backups [Link](https://www.mongodb.com/docs/manual/core/backups/)


