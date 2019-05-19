package main

import (
	"data-tunning-scripts/models"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func GenerateCsvFromSuicideNews() {
	scrapedNewsWithSuicide := models.FindAllFromFilteredCollection()
	accumulator := make(map[string]models.MonthNewsResult)
	for _, result := range scrapedNewsWithSuicide {
		monthCode, month, year := generateMonthCode(result.Date)
		fmt.Println("adding new result on " + monthCode)
		if val, ok := accumulator[monthCode]; ok {
			val.NumberOfSuicideNews = val.NumberOfSuicideNews + 1
			accumulator[monthCode] = val
		} else {
			accumulator[monthCode] = models.MonthNewsResult{MonthCode: monthCode, Month: month, Year: year, NumberOfSuicideNews: 1}
		}
	}
	csv := models.ConvertToCsv(accumulator)
	bytes := []byte(csv)
	ioutil.WriteFile("../data/processed/processed_news_suicide_data/scraped_news_suicide_data.csv", bytes, 0644)

	csv = sortCsv(csv)
	bytes2 := []byte(csv)
	ioutil.WriteFile("../data/processed/processed_news_suicide_data/scraped_news_suicide_data_sorted.csv", bytes2, 0644)
}

func generateMonthCode(date time.Time) (monthCode string, month string, year string) {
	formatedDate := date.Format("2006/01/02")
	month = strings.Split(formatedDate, "/")[1]
	year = strings.Split(formatedDate, "/")[0]
	monthCode = year + month + "00"
	return monthCode, month, year
}

func sortCsv(csv string) string {
	csvSlice := strings.Split(csv, "\n")
	sort.Strings(csvSlice)
	return strings.Join(csvSlice, "\n")
}

func main() {
	GenerateCsvFromSuicideNews()
}
