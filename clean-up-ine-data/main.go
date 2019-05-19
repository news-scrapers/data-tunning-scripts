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

func cleanUpCsv(csvStr string, year string) (csvAll string, csvAmbos string, csvMujeres string, csvHombres string) {
	fmt.Println(year)
	csv := strings.Split(csvStr, "\n")
	csvNoHeading := csv[5:58]
	resultCsvAll := []string{}
	resultCsvAmbos := []string{}
	resultCsvHombres := []string{}
	resultCsvMujeres := []string{}

	currentMonthInLine := 0
	for index, line := range csvNoHeading {
		if index == 0 {
			//line = strings.ReplaceAll(line, "�", "n")
			//newLine := "mes;" + line
			//resultCsvAll = append(resultCsvAll, newLine)
		} else if len(strings.Split(line, ";")) == 2 && index > 4 {
			currentMonthInLine = currentMonthInLine + 1
		} else if index > 4 {
			month := formatMonth(currentMonthInLine)
			line = strings.ReplaceAll(line, "    ", "")
			line = strings.ReplaceAll(line, ".0", "")
			newLine := year + month + "00;" + line
			resultCsvAll = append(resultCsvAll, newLine)

			if strings.Contains(line, "Ambos") {
				resultCsvAmbos = append(resultCsvAmbos, newLine)
			} else if strings.Contains(line, "Muj") {
				resultCsvMujeres = append(resultCsvMujeres, newLine)
			} else {
				resultCsvHombres = append(resultCsvHombres, newLine)
			}
		}
	}
	return strings.Join(resultCsvAll, ""), strings.Join(resultCsvAmbos, ""), strings.Join(resultCsvMujeres, ""), strings.Join(resultCsvHombres, "")
}
func formatMonth(currentMonthInLine int) string {
	var month string
	if currentMonthInLine > 9 {
		month = strconv.Itoa(currentMonthInLine)
	} else {
		month = "0" + strconv.Itoa(currentMonthInLine)
	}
	return month
}
func main() {
	path := "../data/raw_ine_data/all/"
	pathOut := "../data/processed/processed_ine_data/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	header := "mes;descrip;Todas las edades;Menores de 15 a�os;De 15 a 29 a�os;De 30 a 39 a�os;De 40 a 44 a�os;De 45 a 49 a�os;De 50 a 54 a�os;De 55 a 59 a�os;De 60 a 64 a�os;De 65 a 69 a�os;De 70 a 74 a�os;De 75 a 79 a�os;De 80 a 84 a�os;De 85 a 89 a�os;De 90 a 94 a�os;De 95 a�os y m�s;"
	csvFinal, csvAmbosFinal, csvMujeresFinal, csvHombresFinal := header, header, header, header
	for _, f := range files {
		filename := f.Name()
		year := strings.Split(filename, "_")[2]
		fmt.Println(filename)
		rawCsvStr := readAsString(path + filename)
		csv, csvAmbos, csvMujeres, csvHombres := cleanUpCsv(rawCsvStr, year)
		csvFinal = csvFinal + "\n" + csv
		csvAmbosFinal = csvAmbosFinal + "\n" + csvAmbos
		csvMujeresFinal = csvMujeresFinal + "\n" + csvMujeres
		csvHombresFinal = csvHombresFinal + "\n" + csvHombres
	}
	saveStringToFile(csvFinal, pathOut+"procesado_suicidio_edad_todo.csv")
	saveStringToFile(csvAmbosFinal, pathOut+"procesado_suicidio_edad_ambos.csv")
	saveStringToFile(csvMujeresFinal, pathOut+"procesado_suicidio_edad_mujeres.csv")
	saveStringToFile(csvHombresFinal, pathOut+"procesado_suicidio_edad_hombres.csv")

}
