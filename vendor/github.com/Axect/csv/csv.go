package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Write writes csv file
func Write(List [][]string, name string) {
	Title := name
	file, err := os.Create(Title)
	checkError("Cannot create a file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range List {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
	fmt.Println("Complete to Write")
}

// Read read csv file
func Read(directory string) [][]string {
	Title := directory
	file, err := os.Open(Title)
	checkError("Cannot open file", err)

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	return rows
}

// Convert converts [][]string to [][]float64
func Convert(str [][]string) [][]float64 {
	Temp := make([][]float64, len(str), len(str))
	for i := range str {
		for j := range str[i] {
			temp, _ := strconv.ParseFloat(str[i][j], 64)
			Temp[i] = append(Temp[i], temp)
		}
	}
	return Temp
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
