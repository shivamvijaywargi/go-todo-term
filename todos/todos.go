package todos

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func List() {
	file, err := os.Open("todos.csv")
	if err != nil {
		log.Fatal("Error reading the todos csv file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV", err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
