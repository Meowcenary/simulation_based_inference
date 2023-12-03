package main

import (
	"testing"

  "github.com/Meowcenary/simulation_based_inference/boot"
  "github.com/Meowcenary/simulation_based_inference/util"
)

/*
This variable is used to ensure the compiler does not optimize away the call to stats.LinearRegression
See the Benchmarks section in Jon Bodner's Learning Go: An Idiomatic Approach, Chapter 13
*/
var blackhole float64
func BenchmarkBoot(b *testing.B) {
    seed := int64(123)
    random := util.SeedRand(seed)

    data := util.GenerateRandomVector(random, 1000)
    result, _ := boot.Boot(data, seed, util.Mean, 1000)
    blackhole = result
}

// same as above blackhole variable, but for matrix result
var blackholeMatrix []float64
func BenchmarkBootMatrix(b *testing.B) {
    seed := int64(123)
    random := util.SeedRand(seed)

    data := util.GenerateRandomMatrix(random, 1000, 20)
    result, _ := boot.BootMatrix(data, seed, util.MeanMatrix, 1000)
    blackholeMatrix = result
}
