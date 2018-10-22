package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbname    = "homework"
	dbuser    = "postgres"
	maxWorker = 100
)

func main() {

	numWorkers := flag.Int("numWorkers", 1, fmt.Sprintf("Number of parallel connections to the server (min 1, max %d)", maxWorker))
	file := flag.String("file", "", "CSV file with column format \"hostname, starttime, endtime\" with header intact.")

	flag.Parse()

	if *numWorkers < 0 || *numWorkers > maxWorker || *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname)
	db, err := sql.Open("postgres", connStr)

	checkError(err)
	defer db.Close()

	// db.Open does not open the connection. Hence, validate DSN.
	err = db.Ping()

	// check if this validation would work with channels
	checkError(err)

	queries := parseCSV(*file + ".csv")

	work := make(chan Request, len(queries))
	defer close(work)

	done := make(chan struct{})
	defer close(done)

	go request(work, queries, done)

	NewScheduler(*numWorkers, db).schedule(work, done)
}

func request(work chan Request, queries []Query, done chan struct{}) {
	c := make(chan Result)

	results := make([]Result, 0, len(queries))

	for _, query := range queries {
		work <- Request{query, c}
		results = append(results, <-c)
	}

	GetStats(results).print()

	done <- struct{}{}
}
