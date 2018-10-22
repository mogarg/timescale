# Benchmarking

Benchmarks takes a number of workers and a .csv file with format "hostname, start_time, end_time" (header included) to generate and benchmark query run time. Queries are allotted to workers based on hostnames (no two workers share queries that touch same hostnames).

## Build and Usage

```bash
go get github.com/mogarg/timescale

./benchmark -numWorkers 10 -file query_params
```

## Further Additions

1. Testing
2. Generating queries on the fly (currently uses a base query)
3. Balanced Scheduling: Alot next hostname to some worker with the least amount of pending queries to run (could be useful for bigger workloads.
4. More stats by percentiles.
5. Any architectural changes suggested.
6. Better errors handling if some queries fail.