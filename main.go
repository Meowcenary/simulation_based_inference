package main

import (
	"fmt"
  "math/rand"

  "github.com/Meowcenary/simulation_based_inference/boot"
  "github.com/Meowcenary/simulation_based_inference/util"
)

// mean functions
func mean(data []float64) float64 {
  total := 0.0
  for i := 0; i < len(data); i++ {
    total += data[i]
  }
  return total/float64(len(data))
}

// function that determines the mean of each column of a data matrix
func meanMatrix(data [][]float64) []float64 {
  // set totals to the length of the first row
  totals := make([]float64, len(data[0]))

  // for each row of data
  for i := 0; i < len(data); i++ {
    // for each column within a row of data
    for j := 0; j < len(data[i]); j++ {
      totals[j] += data[i][j]
    }
  }

  for i := 0; i < len(totals); i++ {
    totals[i] = totals[i]/float64(len(data))
  }

  return totals
}

func main() {
	// set to same value as seed of main.R
	seed := int64(123)
	source := rand.NewSource(seed)
	random := rand.New(source)

	// data stored as matrix
	dataMatrix := util.GenerateRandomMatrix(random, 100, 2)
	resultMatrix, err := boot.BootMatrix(dataMatrix, seed, meanMatrix, 1000)
  if err != nil {
    panic(err)
  }
	fmt.Println("Matrix data mean result ", resultMatrix)

	// data stored as vector
	dataVector := util.GenerateRandomVector(random, 100)
	resultVector, err := boot.Boot(dataVector, seed, mean, 1000)
  if err != nil {
    panic(err)
  }
	fmt.Println("Vector data mean result ", resultVector)
}
