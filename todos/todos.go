package todos

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Todo struct {
	Id        int    `json:"id"`
	Item      string `json:"item"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

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

func Add() {
	file, err := os.Create("todos.csv")
	if err != nil {
		log.Fatal("Error reading CSV", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	extraData := []string{
		"One",
		"Added from add command",
		"Pending",
		"Today",
	}

	writer.Write(extraData)
}
