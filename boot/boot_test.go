package boot

import (
  "math"
	"math/rand"
  "reflect"
	"testing"

  "github.com/Meowcenary/simulation_based_inference/util"
)

// rand instance for tests
func seed() int64 {
	return 123
}

func random() *rand.Rand {
  return rand.New(rand.NewSource(seed()))
}

func mean(data []float64) float64 {
  total := 0.0
  for i := 0; i < len(data); i++ {
    total += data[i]
  }
  return total/float64(len(data))
}

// tolerance is amount of difference allowable for comparing float64 values
func tolerance() float64 {
  return 0.0000000001
}

// boot tests

// runs without raising errors
func TestBootError(t *testing.T) {
  data := []float64{11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0}

  _, err := Boot(data, seed(), mean, 1000)

  if err != nil {
    t.Errorf("Error raised on call to boot")
  }
}

// returns reasonable value for mean function on resampling
func TestMeanFunction(t *testing.T) {
  // the mean of the set without sampling is 5
  data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
  calculatedMean, err := Boot(data, seed(), mean, 1000)
  expectedMean := float64(5.750000)

  if err != nil {
    t.Errorf("Error raised on call to boot")
  } else if math.Abs(expectedMean-calculatedMean) >= 0.0000001 {
    t.Errorf("The mean returned from boot does not match what is expected for this data and seed:\nExpected: %f\nGot: %f", expectedMean, calculatedMean)
  }
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

// runs without raising errors
func TestBootMatrixError(t *testing.T) {
  data := [][]float64{
    {1.0, 2.0, 3.0, 4.0},
    {5.0, 6.0, 7.0, 8.0},
    {9.0, 10.0, 11.0, 12.0},
  }

  // calculatedMeanResults, err := BootMatrix(data, seed(), meanMatrix, 1000)
  _, err := BootMatrix(data, seed(), meanMatrix, 1000)


  if err != nil {
    t.Errorf("Error raised on call to boot")
  }
}

func TestMeanMatrixFunction(t *testing.T) {
  data := [][]float64{
    {1.0, 2.0, 3.0, 4.0},
    {5.0, 6.0, 7.0, 8.0},
    {9.0, 10.0, 11.0, 12.0},
  }

  // calculatedMeanResults, err := BootMatrix(data, seed(), meanMatrix, 1000)
  calculatedAverages, err := BootMatrix(data, seed(), meanMatrix, 1000)
  expectedAverages := []float64{0.011, 0.014, 0.017, 0.02}
  if err != nil {
    t.Errorf("Error raised on call to boot")
   }else if !reflect.DeepEqual(expectedAverages, calculatedAverages) {
    t.Errorf("Expected: %v\nGot: %v", expectedAverages, calculatedAverages)
  }
}
