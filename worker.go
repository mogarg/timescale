package main

import (
	"database/sql"
)

// Worker executes queries
type Worker struct {
	ID       int
	requests chan Request

	processed int
}

func (w *Worker) work(done chan *Worker, db *sql.DB) {
	for {
		request := <-w.requests
		request.result <- request.runQuery(db)
		done <- w
	}
}
