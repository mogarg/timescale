package main

import (
	"database/sql"
	"time"
)

// Request takes a query and returns the results
type Request struct {
	query  Query
	result chan time.Duration
}

// Query defines the model for building the queries
type Query struct {
	Hostname  string
	Starttime time.Time
	Endtime   time.Time
}

const (
	baseQuery = `SELECT time_bucket('1 minute', ts) as minute,
	MAX(usage) as max_usage,
	MIN(usage) as min_usage
	FROM cpu_usage
	WHERE host = $1 AND ts BETWEEN $2 AND $3 
	GROUP BY minute`
)

func (r *Request) runQuery(db *sql.DB) time.Duration {

	parsed, err := db.Prepare(baseQuery)

	checkError(err)

	query := r.query

	var duration time.Duration

	start := time.Now()

	_, err = parsed.Exec(query.Hostname, query.Starttime, query.Endtime)
	if err == nil {
		duration = time.Now().Sub(start)
	}

	return duration
}
