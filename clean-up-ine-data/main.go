package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readAsString(filename string) string {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	return str
}

func saveStringToFile(str string, filename string) error {
	bytes := []byte(str)
	err := ioutil.WriteFile(filename, bytes, 0644)
	return err
}

func cleanUpCsv(csvStr string, year string) string {
	fmt.Println(year)
	csv := strings.Split(csvStr, "\n")
	csvNoHeading := csv[5:58]
	resultCsv := []string{}
	currentMonthInLine := 0
	for index, line := range csvNoHeading {
		if index == 0 {
			newLine := "mes;" + line
			resultCsv = append(resultCsv, newLine)
		} else if len(strings.Split(line, ";")) == 2 {
			currentMonthInLine = currentMonthInLine + 1
		} else {
			var month string
			if currentMonthInLine > 9 {
				month = strconv.Itoa(currentMonthInLine)
			} else {
				month = "0" + strconv.Itoa(currentMonthInLine)
			}
			newLine := year + month + "00;" + line
			resultCsv = append(resultCsv, newLine)
		}
	}

	return strings.Join(resultCsv, "")
}

func main() {
	path := "../data/all/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filename := f.Name()
		year := strings.Split(filename, "_")[2]
		fmt.Println(filename)
		rawCsvStr := readAsString(path + filename)
		csv := cleanUpCsv(rawCsvStr, year)
		saveStringToFile(csv, "example.csv")
		break
	}

}
