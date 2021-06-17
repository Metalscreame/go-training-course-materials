package main

import (
	"bufio"
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

//go:embed test.csv
var embedded embed.FS

func main() {
	//data, err := embedded.ReadFile("migration.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//reader := csv.NewReader(bytes.NewReader(data))

	csvFile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line[0], line[1])
	}
}


// https://github.com/gocarina/gocsv - custom unmarshall csv lib