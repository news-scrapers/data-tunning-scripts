package main

import (
	"data-tunning-scripts/models"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sort"
	"time"
)

func GenerateCsvFromSuicideNewsWithSentiment() {
	excludedWords := []string{"atentado", "Atentado", "atentar", "Atentar", "terrorista", "Terrorista", "inmola", "Inmola"}

	accumulator := make(map[string]models.MonthNewsResultWithSentiment)
	root := "../data/processed/processed_news_suicide_data/resultsSuicide/"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err!=nil {
			fmt.Println(err)
			fmt.Println(info)
		}
		file, _ := ioutil.ReadFile(path)

		result := models.NewScraped{}
	
		_ = json.Unmarshal([]byte(file), &result)
		if contaisExcluded(result.Content, excludedWords) == false {
		monthCode, month, year := generateMonthCode(result.Date)
			fmt.Println("adding new result on " + monthCode)
			if val, ok := accumulator[monthCode]; ok {
				val.NumberOfSuicideNews = val.NumberOfSuicideNews + 1
				val.AverageSentimentResult = val.AverageSentimentResult + result.SentimentAnalysis.Comparative
				accumulator[monthCode] = val
			} else {
				accumulator[monthCode] = models.MonthNewsResultWithSentiment{MonthCode: monthCode, Month: month, Year: year, NumberOfSuicideNews: 1}
			}
		}
		
        return nil
	})
	csv := models.ConvertToCsvWithSentiment(accumulator)
	bytes := []byte(csv)
	ioutil.WriteFile("../data/processed/processed_news_suicide_data/scraped_news_suicide_data_sentiment.csv", bytes, 0644)

	csv = sortCsv(csv)
	bytes2 := []byte(csv)
	ioutil.WriteFile("../data/processed/processed_news_suicide_data/scraped_news_suicide_data_sentiment_sorted.csv", bytes2, 0644)

	if err!=nil {
		fmt.Println(err)
	}
}


func contaisExcluded(str string, excludedWords []string) bool {
	for _, word := range excludedWords {
		if strings.Contains(str, word) {
			fmt.Println("excluded because contains " + word)
			return true
		}
	}
	return false
}

func generateMonthCode(date time.Time) (monthCode string, month string, year string) {
	formatedDate := date.Format("2006-01-02")
	month = strings.Split(formatedDate, "-")[1]
	year = strings.Split(formatedDate, "-")[0]
	monthCode = year + "-" + month + "-" + "01"
	return monthCode, month, year
}

func sortCsv(csv string) string {
	csvSlice := strings.Split(csv, "\n")
	sort.Strings(csvSlice)
	return strings.Join(csvSlice, "\n")
}



func main() {
	GenerateCsvFromSuicideNewsWithSentiment()
}
