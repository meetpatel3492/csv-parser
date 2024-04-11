package csvparser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsvFile() [][]string {
	file, err := os.Open("./statement.csv")

	if err != nil {
		log.Fatal("Error reading input file ", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Error parsing csv file ", err)
	}

	return records
}

func PrintCsvByColumn(records [][]string){

	var dates []string
	var payee []string
	var amount []string

	for _, record := range records {
		dates = append(dates, record[0])
		payee = append(payee, record[1])
		amount = append(amount, record[2])
	}

	fmt.Println(dates)
	fmt.Println(payee)
	fmt.Println(amount)
}