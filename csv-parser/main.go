package main

import (
	"csv-parser/database"
	"csv-parser/model"
	csvparser "csv-parser/parser"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	database.Connect()
	database.Migrate()
	records := csvparser.ReadCsvFile()

	Date1 := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < len(records); i++ {
		wg.Add(1)
		go func(i int) {
			convertedDate, err := time.Parse("1/2/2006", records[i][0])
			if err != nil {
				log.Fatalf("Error converting date to format - %v, err - %v", time.RFC3339, err)
			}
			amexTransaction := model.AmexTransaction{
				Date: convertedDate,
				Payee: records[i][1],
				Amount: records[i][2],
			}
			log.Println(amexTransaction)
			database.Instance.Create(&amexTransaction)
			wg.Done()
		}(i)
	}
	wg.Wait()
	Date2 := time.Now()
	timeDiff := Date2.Sub(Date1)
	fmt.Println(timeDiff)
}
