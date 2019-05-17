package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDb() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbAddress := os.Getenv("database_url")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	db = client.Database("news-scraped")
}

func GetDb() *mongo.Database {
	return db
}
