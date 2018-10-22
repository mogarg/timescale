# Benchmarking

A tool for benchmarking TimeScaleDB

Benchmarks takes the number of workers and a .csv file (.csv removed) with format "hostname, start_time, end_time" (header included) to generate and benchmark query run time. Queries are allotted to workers based on hostnames (no two workers share queries that touch the same hostnames).

## Build and Usage

```bash
go get github.com/mogarg/timescale

./benchmark -numWorkers 10 -file query_params
```

## Further Additions

-[ ] Testing
-[ ] Generating queries on the fly (currently uses a base query)
-[ ] Balanced Scheduling: Alot next hostname to some worker with the least amount of pending queries to run (could be useful for bigger workloads.
-[ ] More stats by percentiles.
-[ ] Any architectural changes suggested.
-[ ] Better errors handling if some queries fail.
