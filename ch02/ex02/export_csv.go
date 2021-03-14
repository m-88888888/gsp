package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"No.", "taskname", "status"},
		{"1", "shopping", "WIP"},
		{"2", "studying", "TODO"},
		{"3", "lunch with Boby", "DONE"},
	}

	if err := WriteCsvFile(records); err != nil {
		log.Fatal(err)
	}
}

func WriteCsvFile(records [][]string) error {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return nil
}
