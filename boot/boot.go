package boot

import (
	"fmt"
	"math/rand"

  "github.com/Meowcenary/simulation_based_inference/util"
  // "reflect"
)

// Error to catch data that is not supported by Boot
type AttributeError struct {}

func (a *AttributeError) Error() string {
  return "Error with attribute passed to Boot"
}

// Note to self - the vector and matrix resample methods should return the same results if the data
// passed in the matrix version is one column i.e a vertical array and if the seed is the same

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
// seems like you will need to split into boot vector and boot matrix
func Boot(data []float64, seed int64, statistic func([]float64) float64, r int) (float64, error) {
  // TODO add check for if it should be run concurrently
  stype := "i"
  source := rand.NewSource(seed)
  random := rand.New(source)
  call := fmt.Sprintf("Boot(data: [%f,..., %f], seed: %d, random: %v, statistic: %v, r: %d stype: %q)",
		data[0], data[1], seed, random, util.GetFunctionName(statistic), r, stype)

  samples := createBootstrapSamplesVector(random, []float64(data), r)

  // with the samples created run "statistic" function on sets
  // statistic(data)
  total := 0.0
  for i := 0; i < len(samples); i++ {
    total += statistic(samples[i])
  }
  fmt.Println(call)
  // need to cast to float64 for the division to work
  return (total/float64(r)), nil
}

/*
returns []float64 where each value is the average result of statistic() for each column from the bootstrap samples

*/
func BootMatrix(data [][]float64, seed int64, statistic func([][]float64) []float64, r int) ([]float64, error) {
  source := rand.NewSource(seed)
  random := rand.New(source)

  samples := createBootstrapSamplesMatrix(random, [][]float64(data), r)

  // totals is []float64 that is length of a row of data. All values are 0.0
  totals := make([]float64, len(data[0]))

  // for each bootstrap sample...
  for i := 0; i < len(samples); i++ {
    // calculate the statistic for the sample
    results := statistic(samples[i])

    // for each column within the results...
    for j := 0; j < len(totals); j++ {
      // add to the total for the calculation
      totals[j] += results[j]
    }
  }

  for i := 0; i < len(totals); i++ {
    totals[i] = totals[i]/float64(r)
  }

  return totals, nil
}

/*
private method to create bootstrap samples for a vector, returns an array of float64 arrays

r - the instance of rand to generate the samples with
data - the data to be sampled as a two dimensional array of float64 values where each
       row is a multivariate value
samples - r renamed for personal clarity, number of samples to take
*/
func createBootstrapSamplesVector(r *rand.Rand, data []float64, samples int) [][]float64 {
  initialSample := createBootstrapSampleVector(r, data)
  bootstrapSamples := make([][]float64, samples)

  for i := 0; i < samples; i++ {
    bootstrapSamples[i] = createBootstrapSampleVector(r, initialSample)
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

/*
create bootstrap samples for data formatted as a matrix

params:
r -
data -
samples - the number of samples to create (this is "r" renamed for personal clarity)
*/
func createBootstrapSamplesMatrix(r *rand.Rand, data[][]float64, samples int) [][][]float64 {
  sampleSize := len(data)
  initialSample := createBootstrapSampleMatrix(r, data)
  bootstrapSamples := make([][][]float64, sampleSize)

  for i := 0; i < sampleSize; i++ {
    bootstrapSamples[i] = createBootstrapSampleMatrix(r, initialSample)
  }

  return bootstrapSamples
}

/*
in matrix format each row of the matrix represents a
row of data where each index of the row acts as a column

params:
r -
data -

returns:
[][]float64 matrix that has values randomly selected from the data passed to the function
*/
func createBootstrapSampleMatrix(r *rand.Rand, data [][]float64) [][]float64 {
  sampleSize := len(data)
  bootstrapSample := make([][]float64, sampleSize)

  for i := 0; i < sampleSize; i++ {
    bootstrapSample[i] = data[r.Intn(sampleSize)]
  }

  return bootstrapSample
}
