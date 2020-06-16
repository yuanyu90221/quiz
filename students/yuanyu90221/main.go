package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// open the file
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of  'question, answer' ")
	flag.Parse()
	csvfile, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// Iteracte through the records
	correct := 0
	wrong := 0
	idx := 0
	for {
		// Read each problem set from csv
		record, err := r.Read()
		if err == io.EOF {
			fmt.Printf("Result correct rate: %d out of %d\n", correct, idx)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		idx++
		fmt.Printf("Question %d: %s? ", idx, record[0])
		solution, errFormat := strconv.Atoi(record[1])
		if errFormat != nil {
			log.Fatal(errFormat)
		}
		var answer int
		fmt.Scanf("%d", &answer)
		if solution == answer {
			correct++
		} else {
			wrong++
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
