package boot

import (
	"fmt"
	"math/rand"
  "reflect"
)

// Note to self - the vector and matrix resample methods should return the same results if the data
// passed in the matrix version is one column i.e a vertical array and if the seed is the same

// Error to catch data that is not supported by Boot
type AttributeError struct {}

/*
public function for boostrapping, returns the result of the statistical function for the data sets

params:
data - the data to be used for bootstrapping. Currently supports vector and matrix
seed - the seed provided to rand for bootstrapping
random - pointer to an instance of rand created from seed
statistic - function to be run on the data set - in it's current state this rigid but interfaces may
						be an option for generalizing function signatures allowed
r - the number of bootstrap samples to be created
call - the original call to boot (i.e the arguments used to construct the boot object)
stype - string indicating what the second argument of statistic represents - only plan to implement "i" indices to start
				but "f" frequencies and "w" weights are implemented in the R package

returns:
float64, error - float64 is result of statistical function passed as argument, float64 is nil if error
*/
func Boot(data interface{}, seed int, statistic func([]float64) float64, r int) (float64, error) {
  // TODO add check for if it should be run concurrently
	stype = "i"
  call := fmt.Sprintf("Boot(data: %v, seed: %d, random: %v, statistic: %v, r: %i stype: %q)",
		data, seed, random, statistic, r, stype)

  // determine if data is a matrix or vector and call the appropriate function
  dataFormat := reflect.TypeOf()
  if dataFormat == [][]float64 {
    createBootstrapSamplesMatrix(rand.New(rand.Seed(seed)), data, r)
  } else if dataFormat == []float64 {
	  createBootstrapSamplesVector(rand.New(rand.Seed(seed)), data, r)
  } else {
    return nil, AttributeError
  }

  // with the samples created run "statistic" function on sets
  // statistic(data)
}

/*
private method to create bootstrap samples for a vector, returns an array of float64 arrays

r - the instance of rand to generate the samples with
data - the data to be sampled as a two dimensional array of float64 values where each
       row is a multivariate value
samples - r renamed for personal clarity
*/
func createBootstrapSamplesVector(r *rand.Rand, data []float64, samples int) [][]float64 {
  initialSample := createBootstrapSampleVector(samples, data)
  bootstrapSamples := make([][]Point, samples)

  for i := 0; i < samples; i++ {
    bootstrapSamples[i] = createBootstrapSampleVector(r, data)
  }

  return bootstrapSamples
}

func createBootstrapSampleVector(r *rand.Rand, data []float64) []float64 {
  sampleSize := len(data)
  bootstrapSample := make([]float64, sampleSize)

  for i := 0; i < sampleSize; i++ {
    bootstrapSample[i] = data[r.Intn(sampleSize)]
  }

  return bootstrapSample
}

// private method to create samples for Boot function from matrix,
// r - the instance of rand to generate the samples with
// data - the data to be sampled as a two dimensional array of float64 values where each
//        row is a multivariate value
// samples - r renamed for personal clarity
func createBootstrapSamplesMatrix(r *rand.Rand, data [][]float64, samples int) [][]float64 {
  // the initial sample from which the bootstrap samples will to return are drawn
  initialSample := createBootstrapSampleMatrix(r, data)

  // number of bootstrap samples to create
  samples := 1000
  bootstrapSamples := make([][]Point, samples)
  for i := 0; i < samples; i++ {
    // initialSample is now selected from instead of the original data
    bootstrapSamples[i] = createBootstrapSampleMatrix(r, initialSample)
  }

  return bootstrapSamples
}

// the first argument is a pointer to rand.Rand so that the same random seed can be used all tests
// there is probably a better way to do this, but given time constraints this was the approach I went with
func createBootstrapSampleMatrix(r *rand.Rand, data [][]float64) [][]float64 {
  sampleSize := len(data)
  bootstrapSample := make([]Point, sampleSize)

  for i := 0; i < sampleSize; i++ {
    bootstrapSample[i] = data[r.Intn(sampleSize)]
  }

  return bootstrapSample
}
