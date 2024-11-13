package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"USA", "SanJose", "California"},
		{"USA", "Dallas", "Texas"},
		{"India", "Mumbai", "Maharashtra"},
	}
	csvfile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	osWriter := csv.NewWriter(csvfile)
	if err := osWriter.WriteAll(rows); err != nil {
		log.Fatal(err)

		return
	}
	defer csvfile.Close()
}
