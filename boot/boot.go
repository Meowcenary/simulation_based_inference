package boot

import (
	"fmt"
	"math/rand"
  "reflect"
  "runtime"
  "strings"
)

// Error to catch data that is not supported by Boot
type AttributeError struct {}

func (a *AttributeError) Error() string {
  return "Error with attribute passed to Boot"
}

func getFunctionName(temp interface{}) string {
  strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
  return strs[len(strs)-1]
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
  call := fmt.Sprintf("Boot(data: %v, seed: %d, random: %v, statistic: %v, r: %d stype: %q)",
		data, seed, random, getFunctionName(statistic), r, stype)

  // determine if data is a matrix or vector and call the appropriate function
  // switch reflect.TypeOf(data) {
  // // if dataFormat == reflect.TypeOf([][]float64{}) {
  // //   createBootstrapSamplesMatrix(random, data.([][]float64), r)
  // case reflect.TypeOf([]float64{}):
  samples := createBootstrapSamplesVector(random, []float64(data), r)
  // default:
  //   return 0.0, &AttributeError{}
  // }

  // with the samples created run "statistic" function on sets
  // statistic(data)
  total := 0.0
  for i := 0; i < len(samples); i++ {
    total += statistic(samples[i])
  }
  fmt.Println(call)
  // need to cast to float64 for the division to work
  return (total/float64(len(samples))), nil
}

/*
private method to create bootstrap samples for a vector, returns an array of float64 arrays

r - the instance of rand to generate the samples with
data - the data to be sampled as a two dimensional array of float64 values where each
       row is a multivariate value
samples - r renamed for personal clarity
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

// func genericBootstrapSample[]()
