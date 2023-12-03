# Simulation Based Inference
Limited implementation of the method "boot" from R's "boot" package in Go

### Overview
The bulk of the Go code is in in the boot package which itself is under the
directory "boot/". The tests within "boot/boot_test.go" show examples of how the
individual functions work.

At this point in time there are two separate functions to handle the data
formats supported by R's "boot" package, BootVector for single variable data
and BootMarix for multivariate data. In the R package data stored in a matrix is
processed such that row of the matrix is a single multivariate observation. A
potentially useful refactor would be to use generics to create a single function
Boot that could support multiple data types, but given time constraints it made
the most sense to simply implement separate functions.

The R code is all stored under the directory "rscripts/". Each script is a
separate example of using R's boot package.

### Comparing Implementations
One of the primary advantages of R is the relative ease of use. Two of the major
contributing factors to this are the incredible amount of packages available
through The Comprehensive R Archive Network ([CRAN](https://cran.r-project.org/)) and it's use of dynamic typing.
Conversely, Go has significantly less packages available and is statically typed
thus requiring a bit more programming knowledge to work effectively with. The
benefit of building equivalent tools in Go is an incredible increase to
performance. The benchmarks used for testing had each implementation create a
data set of a 1000 value vector and a 1000x2 matrix and then take the average of
the values stored within. Each benchmark was run 100 times and on average Go was
able to process the matrix in 0.02956 nanoseconds and the vector in 0.01018
nanoseconds while R was able to process the matrix in 0.06628799 seconds
(66287990 ns) and the vector in 0.03188 seconds (31880000 ns). Though this
remarkable increase in performance comes at the cost of requiring higher
technical knowledge to work within the more constrained Go environment, it seems
well worth investigating if performance is a concern.

As alluded to in the overview, the biggest challenge in implementing the Go
version of the code was handling data structured in different formats such as
vector or matrix. A significant amount of time was spent attempting to use
generics, a relatively new feature in Go, to create a more general boostrapping
function, but it proved too challenging for the scope of this project. Instead
separate functions were written to support matrix and vector data formats which
led to significantly more code, but was relatively simple to implement and
understand.

Given time constraints there were addtional features in the original boot
function in R that there was not time to implement, but these were largely
optional features for more advanced use cases. The full manual for R's "boot"
package is linked at the bottom of this README.

### Installation
- HTTP: `git clone https://github.com/Meowcenary/simulation_based_inference.git`
- SSH: `git clone git@github.com:Meowcenary/simulation_based_inference.git`

### Running the Program
The file "main.R" imports the package "boot" and then uses the bootstrap method
it provideds to create sample sets from which the mean of the sets is
calculated. The file "main.go" also does this, but using the Go implementation
from the package "boot" defined in this repository.

### Running the Tests
To run all the tests use `go test ./...` from the root directory of the project.
Alternatively use `go test ./<package_name>` to run an individual package's
tests. For example to run the tests for the package `boot`:
    `go test ./boot`

### Running the Benchmarks
To run the Go benchmarks use this command from the project root:
```
go test -bench=.
```

To run the R benchmarks use this command from the project root:
```
Rscript rscripts/benchmark_boot.R
```

Included in the project is a simple bash script to repeatedly run the benchmarks
that can be run with this command from the project root:
```
bash runbenchmarks.sh
```

### R "boot" Package Information
- [Boot package on CRAN](https://cran.r-project.org/web/packages/boot/)
- [Boot PDF manual](https://cran.r-project.org/web/packages/boot/boot.pdf)
