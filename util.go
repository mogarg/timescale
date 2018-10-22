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
			fmt.Printf("record on line %v: incorrect time format\n", line)
			continue
		}

		queries = append(queries, Query{
			Hostname:  line[0],
			Starttime: startime,
			Endtime:   endtime,
		})
	}

	return
}

// Map applies function f over an array of res to convert it to result timing
func Map(res []Result, f func(time.Duration) float64) []ResultTiming {
	resm := make([]ResultTiming, len(res))

	for i := range res {
		resm[i] = ResultTiming{f(res[i].duration), res[i].count}
	}

	return resm
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
