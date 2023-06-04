package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/domain/user/dto"
	"main/domain/user/entity"
)

type Repository struct {
	client *mongo.Collection
}

func NewUserModel(ur *mongo.Collection) *Repository {
	return &Repository{ur}
}

func (r *Repository) findAll() []entity.User {
	cur, err := r.client.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var results []entity.User
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
	fmt.Println(results)
	return results
}

func (r *Repository) findOne(dtoFindOne dto.FindOne) entity.User {
	var results entity.User
	objID, _ := primitive.ObjectIDFromHex(dtoFindOne.Id)
	filter := bson.D{{"_id", objID}}
	err := r.client.FindOne(context.TODO(), filter).Decode(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(dtoFindOne.Id, dtoFindOne, results, filter)
	return results
}
