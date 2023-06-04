package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	host   = "mongodb://127.0.0.1:27017"
	dbName = "test"
	ctx    = context.TODO()
)

var dbCon *mongo.Database

func connection() *mongo.Database {
	clientOption := options.Client().ApplyURI(host)
	client, err := mongo.NewClient(clientOption)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connections")
	dbCon = client.Database(dbName)

	return dbCon
}

func Start() *mongo.Database {
	if dbCon != nil {
		return dbCon
	}
	return connection()
}

func GetCollection(col string) *mongo.Collection {
	if dbCon == nil {
		log.Fatal("MongoDB is not started yet!")
	}
	retCol := dbCon.Collection(col)
	return retCol
}
