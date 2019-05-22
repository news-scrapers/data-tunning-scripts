package main

import (
	"data-tunning-scripts/models"
	"fmt"
)

func main() {
	suicideWords := []string{"suicidio", "suicidar", "suicidando", "suicidios", "suicida", "suicidas", "suicidó", "suicidaron", "suicidado", "suicidando", "suicidados", "suicidaran", "suicidé", "suicide", "suicidaste", "suicidará", "suicidarán", "suicidarían"}
	//suicideWords := []string{"suicidio"}

	newsMap := map[string]models.NewScraped{}

	for _, word := range suicideWords {
		newsScraped := models.SearchNewsWithText(word)
		for _, result := range newsScraped {
			newsMap[result.Url] = result
		}
	}

	resultsWithoutRepetition := []models.NewScraped{}
	for _, result := range newsMap {
		resultsWithoutRepetition = append(resultsWithoutRepetition, result)
	}
	fmt.Println("")
	fmt.Printf("Saving %v news to file", len(resultsWithoutRepetition))

	searchResults := models.SearchResults{}
	searchResults.NewsScrapedResults = resultsWithoutRepetition
	searchResults.SaveToFile("resultsSuicide.json")

	fmt.Println("")
	fmt.Printf("Saving %v news to db", len(resultsWithoutRepetition))
	models.CreateManyNewsScraped(resultsWithoutRepetition)

}
