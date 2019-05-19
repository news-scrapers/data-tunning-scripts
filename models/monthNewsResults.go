package models

import "strconv"

type MonthNewsResult struct {
	MonthCode           string `json:"month_code"`
	Month               string `json:"month"`
	Year                string `json:"year"`
	NumberOfSuicideNews int    `number_of_suicide_news`
}

func ConvertToCsv(results map[string]MonthNewsResult) (csv string) {
	header := "month_code;month; year;number_of_suicide_news"
	csv = header
	for _, result := range results {
		csv = csv + "\n" + result.MonthCode + ";" + result.Month + ";" + result.Year + ";" + strconv.Itoa(result.NumberOfSuicideNews)
	}
	return csv
}
