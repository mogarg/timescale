package main

import (
	"fmt"
	"sort"
	"time"
)

// Stats encapsulates the stats for a list of time.Duration values
type Stats struct {
	Count   int
	Total   ResultTiming
	Average float64
	Median  float64
	Minimum ResultTiming
	Maximum ResultTiming
}

// ResultTiming represents the time spent executing
// a query and the number of results generated for the query
type ResultTiming struct {
	time    float64
	results int
}

func (rt *ResultTiming) stringify() string {
	return fmt.Sprintf("%.2f ms (%d results)", rt.time, rt.results)
}

func (s *Stats) print() {

	fmt.Println("")
	fmt.Printf("Total queries:%d \n", s.Count)
	fmt.Printf("Total time: %s\n", s.Total.stringify())
	fmt.Printf("Average time: %.2f ms\n", s.Average)
	fmt.Printf("Median Time:  %.2f ms\n", s.Median)
	fmt.Printf("Minimum time: %s\n", s.Minimum.stringify())
	fmt.Printf("Maximum time: %s\n", s.Maximum.stringify())
	fmt.Println("")

}

// GetStats takes a list of Result values, creates and returns the Stat
func GetStats(results []Result) *Stats {
	stats := &Stats{}
	var n int

	durationToNS := func(d time.Duration) float64 {
		return float64(d.Nanoseconds()) / float64(1000000)
	}

	resTimes := Map(results, durationToNS)

	sort.Slice(resTimes, func(i, j int) bool {
		return resTimes[i].time < resTimes[j].time
	})

	n = len(resTimes)

	if n == 0 {
		return stats
	}

	stats.Count = n
	stats.Minimum = resTimes[0]
	stats.Maximum = resTimes[n-1]

	totalTime := 0.0
	totalRes := 0

	for _, rt := range resTimes {
		totalTime += rt.time
		totalRes += rt.results
	}

	stats.Total = ResultTiming{totalTime, totalRes}

	stats.Average = totalTime / float64(len(resTimes))

	if n%2 == 0 && n > 1 {
		stats.Median = (resTimes[n/2].time + resTimes[n/2+1].time) / 2
	} else {
		stats.Median = resTimes[n/2].time
	}

	return stats
}
