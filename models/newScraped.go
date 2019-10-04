package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type NewScraped struct {
	Page      int       `json:"page"`
	FullPage  bool      `json:"full_page"`
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
	ID        string    `json:"id"`
	SentimentAnalysis Sentiment `json:"sentiment_analysis"`
}

type Sentiment struct {
	Comparative        float64    `json:"comparative"`
}

type SearchResults struct {
	NewsScrapedResults []NewScraped `json:"news_scraped_result"`
}

func FindAllFromFilteredCollection() []NewScraped {
	ConnectDb()
	db := GetDb()

	fmt.Println("searching all news ")
	collection := db.Collection("FilteredNewsContentScraped")
	// $text: { $search: "suicidio" }
	results := []NewScraped{}
	cur, err := collection.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		fmt.Println("error found")
		panic(err)
	}
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewScraped
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error found")
			fmt.Println(err)
		}
		results = append(results, elem)
	}

	fmt.Printf("found %v", len(results))
	return results
}

func SearchNewsWithText(text string) []NewScraped {
	ConnectDb()
	db := GetDb()

	fmt.Println("searching news with text -->" + text)
	collection := db.Collection("NewsContentScraped")
	// $text: { $search: "suicidio" }
	results := []NewScraped{}
	cur, err := collection.Find(context.Background(), bson.M{"$text": bson.M{"$search": text}}, nil)
	if err != nil {
		fmt.Println("error found")
		panic(err)
	}
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewScraped
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error found")
			fmt.Println(err)
		}
		results = append(results, elem)
	}

	fmt.Printf("found %v", len(results))
	return results
}

func SearchNewsWithTextAvoiding(text string, avoid []string) []NewScraped {
	ConnectDb()
	db := GetDb()

	fmt.Println("searching news with text -->" + text)
	collection := db.Collection("NewsContentScraped")
	// $text: { $search: "suicidio" }
	results := []NewScraped{}

	avoidSentence := ""
	for _, word := range avoid {
		avoidSentence = avoidSentence + "-" + word + " "
	}
	fmt.Println(avoidSentence)
	cur, err := collection.Find(context.Background(), bson.M{"$text": bson.M{"$search": text + " " + avoidSentence}}, nil)
	if err != nil {
		fmt.Println("error found")
		panic(err)
	}
	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem NewScraped
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error found")
			fmt.Println(err)
		}
		results = append(results, elem)
	}

	fmt.Printf("found %v", len(results))
	return results
}

func CreateManyNewsScraped(newsScraped []NewScraped) error {
	ConnectDb()

	db := GetDb()
	collection := db.Collection("FilteredNewsContentScraped")
	docs := []interface{}{}

	for _, result := range newsScraped {
		docs = append(docs, result)
	}
	_, err := collection.InsertMany(context.Background(), docs)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (searchResults *SearchResults) SaveToFile(path string) {
	file, errMarshal := json.MarshalIndent(searchResults, "", " ")
	if errMarshal != nil {
		panic(errMarshal)
	}
	errSaving := ioutil.WriteFile(path, file, 0644)
	if errSaving != nil {
		if errMarshal != nil {
			panic(errSaving)
		}
	}
}
