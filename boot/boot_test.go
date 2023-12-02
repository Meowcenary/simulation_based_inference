package boot

import (
  "math"
	"math/rand"
	"testing"
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
