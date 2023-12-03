package util

import (
  "math/rand"
  "reflect"
  "runtime"
  "strings"
)

/*
public function that returns pointer to instance of Rand with specified seed value

params:
seed - int64 value for seeding Rand

returns:
pointer to instance of Rand
*/
func SeedRand(seed int64) *rand.Rand {
	// set to same value as seed of main.R
	source := rand.NewSource(seed)
	return rand.New(source)
}

/*
public function for stringifying a function name. In this repo it is used as to get the name of the function passed
as the argument "statistic" to the two Boot functions (BootVector, BootMatrix) defined in the boot package

params:
function - an interace is used to genericize the function, but the value should be a pointer to a function

returns:
the name of the function as a string
*/
func GetFunctionName(function interface{}) string {
  strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()), ".")
  return strs[len(strs)-1]
}

/*
public function for generating a random vector (Golang slice) of float64 values for use by boot

params:
r - pointer to an instance of rand
size - integer size of the vector to be generated

returns:
slice of float64 values that is size values long
*/
func GenerateRandomVector(r *rand.Rand, size int) []float64 {
  vector := make([]float64, size)

  for i := 0; i < size; i++ {
    vector[i] = rand.NormFloat64()
  }

  return vector
}

/*
public function for generating a rows by by columns matrix (Golang 2d slice) of float64 values for use by boot

params:
r - pointer to an instance of rand
rows - number of rows for the matrix to be generated
columns - number of columns for the matrix to be generated

returns:
matrix (2d slice) of float64 values that has the specified number of rows and columns
*/
func GenerateRandomMatrix(r *rand.Rand, rows, columns int) [][]float64 {
  // create an empty matrix to hold values
  matrix := make([][]float64, rows)
  for i := 0; i < rows; i++ {
    matrix[i] = make([]float64, columns)
  }

  // populate matrix
  for i := 0; i < rows; i++ {
    for j := 0; j < columns; j++ {
      matrix[i][j] = rand.NormFloat64()
    }
  }

  return matrix
}

// mean functions
func Mean(data []float64) float64 {
  total := 0.0
  for i := 0; i < len(data); i++ {
    total += data[i]
  }
  return total/float64(len(data))
}

// function that determines the mean of each column of a data matrix
func MeanMatrix(data [][]float64) []float64 {
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
