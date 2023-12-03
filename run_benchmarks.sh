#!/usr/bin/env bash

# run Go benchmarks 100 times
for ((i = 0 ; i < 100 ; i++)); do
  go test -bench=. | grep ns/op
done

# run R benchmarks 100 times
for ((i = 0 ; i < 100 ; i++)); do
  Rscript rscripts/benchmark_boot.R
done
