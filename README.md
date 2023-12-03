# Simulation Based Inference
Limited implementation of the method "boot" from R's "boot" package in Go

### Overview
The bulk of the Go code is in in the boot package which itself is under the
directory "boot/". The tests within "boot/boot_test.go" show examples of how the
individual functions work.

At this point in time there are two separate functions to handle the data
formats supported by R's "boot" package - BootVector for single variable data
and BootMarix for multivariate data. In the R package data stored in a matrix is
processed such that row of the matrix is a single multivariate observation. A
potentially useful refactor would be to use generics to create a single function
Boot that could support multiple data types, but given time constraints it made
the most sense to simply implement separate functions.

### Installation
- HTTP: `https://github.com/Meowcenary/simulation_based_inference.git`
- SSH: `git@github.com:Meowcenary/simulation_based_inference.git`

### Running The Program
The file "main.R" imports the package "boot" and then uses the bootstrap method
it provideds to create sample sets from which the mean of the sets is
calculated. The file "main.go" also does this, but using the Go implementation
from the package "boot" defined in this repository.

### Running The Tests
To run all the tests use `go test ./...` from the root directory of the project.
Alternatively use `go test ./<package_name>` to run an individual package's
tests. For example to run the tests for the package `boot`:
    `go test ./boot`

### R "boot" Package Information
- [Boot package on CRAN](https://cran.r-project.org/web/packages/boot/)
- [Boot PDF manual](https://cran.r-project.org/web/packages/boot/boot.pdf)
