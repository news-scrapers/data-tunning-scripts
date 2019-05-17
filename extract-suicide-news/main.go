package main

import (
	"context"
	"data-tunning-scripts/models"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var db *mongo.Database

func connectDb() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbAddress := os.Getenv("database_url")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	db = client.Database("news-scraped")
}

func extractNews() {
	fmt.Println("-------")
	collection := db.Collection("NewsContentScraped")
	// $text: { $search: "suicidio" }
	results := []models.NewScraped{}
	cur, err := collection.Find(context.Background(), bson.M{"$text": bson.M{"$search": "suicidio"}}, nil)
	if err != nil {
		fmt.Println("error found")
		panic(err)
	}
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem models.NewScraped
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error found")
			fmt.Println(err)
		}
		results = append(results, elem)
	}

	fmt.Println(len(results))
}

func main() {
	connectDb()
	extractNews()

}
