package entity

import (
	"go.mongodb.org/mongo-driver/mongo"
	adapter "main/adapter/mongo"
)

var collectionName = "user"

type User struct {
	Id    string `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Senha string `json:"senha" bson:"senha"`
}

func Repository() *mongo.Collection {
	return adapter.GetCollection(collectionName)
}
