package main

import (
	"fmt"

  "github.com/Meowcenary/simulation_based_inference/boot"
  "github.com/Meowcenary/simulation_based_inference/util"
)

func main() {
  // instantiate random generator for creating data
  seed := int64(123)
  random := util.SeedRand(seed)

	// data stored as matrix
	dataMatrix := util.GenerateRandomMatrix(random, 100, 2)
	resultMatrix, err := boot.BootMatrix(dataMatrix, seed, util.MeanMatrix, 1000)
  if err != nil {
    panic(err)
  }
	fmt.Println("Matrix data mean result ", resultMatrix)

	// data stored as vector
	dataVector := util.GenerateRandomVector(random, 100)
	resultVector, err := boot.Boot(dataVector, seed, util.Mean, 1000)
  if err != nil {
    panic(err)
  }
	fmt.Println("Vector data mean result ", resultVector)
}
