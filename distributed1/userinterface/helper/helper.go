package helper

import (
	"context"
	"fmt"
	"log"
	"userinterface/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	fmt.Println("Data base storage")
	clientopt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientopt)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("Files").Collection("Meta")
}

func Insert_data(t model.Meta) {
	insert, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insert)
}

func Get_data(t string) (model.Meta, bool) {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var k model.Meta
	for curr.Next(context.Background()) {
		var v model.Meta
		err := curr.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		if v.Name == t {
			return v, true
		}
	}
	return k, false
}
