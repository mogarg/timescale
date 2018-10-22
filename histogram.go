package main

import (
	"fmt"
	"sort"
	"time"
)

// Histogram is a representation of statistics
type Stats struct {
	Count   int
	Minimum float64
	Average float64
	Maximum float64
	Median  float64
}

//
func (s *Stats) printStats() {

	fmt.Println("\n---------Stats--------")
	fmt.Printf("Total Queries: %d\n", s.Count)
	fmt.Printf("Minimum time:  %.2f ms\n", s.Minimum)
	fmt.Printf("Maximum time:  %.2f ms\n", s.Maximum)
	fmt.Printf("Average time:  %.2f ms\n", s.Average)
	fmt.Printf("Median Time:   %.2f ms\n", s.Median)
	fmt.Println("----------------------")

}

// HistogramFromDurations returns a new statsogram for the durations
func StatsFromDurations(durations []time.Duration) *Stats {
	nanos := make([]float64, len(durations))

	for i, d := range durations {
		nanos[i] = float64(d.Nanoseconds()) / float64(1000000)
	}

	stats := &Stats{}

	n := len(nanos)

	if n == 0 {
		return stats
	}

	sort.Float64s(nanos)

	stats.Count = n
	stats.Minimum = nanos[0]
	stats.Maximum = nanos[n-1]

	stats.Average = nanos[0]
	for _, x := range nanos {
		stats.Average += x
	}
	stats.Average /= float64(len(nanos))

	if n%2 == 0 {
		stats.Median = (nanos[n/2] + nanos[n/2+1]) / 2
	} else {
		stats.Median = nanos[n/2]
	}

	return stats
}
