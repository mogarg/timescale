package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	timelayout = "2006-01-02 15:04:05"
)

func parseCSV(filename string) (queries []Query) {

	csvFile, err := os.Open(filename)
	checkError(err)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	checkError(err)

	lines, err := reader.ReadAll()
	checkError(err)

	for i, line := range lines {

		if i == 0 {
			continue
		}

		if err == io.EOF {
			break
		} else {
			checkError(err)
		}

		startime, err := time.Parse(timelayout, line[1])
		endtime, err := time.Parse(timelayout, line[2])
		if err != nil {
			fmt.Print(line)
		}
		checkError(err)

		queries = append(queries, Query{
			Hostname:  line[0],
			Starttime: startime,
			Endtime:   endtime,
		})
	}

	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
