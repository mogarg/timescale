package main

import (
	"database/sql"
)

// Scheduler assigns queries to worker threads
// based on hostnames in the queries. Potential
// for more balanced scheduling if hostname affinity
// is not a required characteristic.
type Scheduler struct {
	pool    Pool
	done    chan *Worker
	i       int
	hostmap map[string]int
}

// Pool is a collection of worker threads
type Pool []*Worker

// NewScheduler returns a scheduler by creating a pool of threads
func NewScheduler(workerCount int, db *sql.DB) *Scheduler {
	done := make(chan *Worker, workerCount)

	b := &Scheduler{make(Pool, 0, workerCount), done, 0, make(map[string]int)}

	for i := 0; i < workerCount; i++ {
		w := &Worker{ID: i, requests: make(chan Request, 1)}
		go w.work(b.done, db)
		b.pool = append(b.pool, w)
	}
	return b
}

func (b *Scheduler) schedule(work chan Request, done chan struct{}) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case _ = <-b.done:
		case <-done:
			b.print()
			return
		}
	}
}

func (b *Scheduler) print() {
	sum := 0

	//	fmt.Printf("\nWork Schedule:\n")

	for _, w := range b.pool {
		//		fmt.Printf("%d ", w.processed)
		sum += w.processed
	}

	_ = float64(sum) / float64(len(b.pool))
	//	fmt.Printf("queries/worker %.2f\n", avg)
}

func (b *Scheduler) dispatch(req Request) {
	var w *Worker

	host := req.query.Hostname

	val, ok := b.hostmap[host]

	if ok == true {
		w = &*b.pool[val]
	} else {
		b.hostmap[host] = b.i
		w = &*b.pool[b.i]
		b.i++
	}

	w.requests <- req
	w.processed++
	if b.i >= len(b.pool) {
		b.i = 0
	}
}
