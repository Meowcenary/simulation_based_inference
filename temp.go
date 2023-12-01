package main

import "fmt"

// Example struct with an interface{} field
type MyStruct struct {
	MyFunction interface{}
}

func SomeFunction(data []float64) float64 {
	// Example implementation of a function
	// You can replace this with your actual implementation
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func main() {
	// Create an instance of MyStruct with the function field assigned
	myInstance := MyStruct{
		MyFunction: SomeFunction,
	}

	// Extract the function from MyStruct and use it
	myFunction, ok := myInstance.MyFunction.(func([]float64) float64)
	if !ok {
		fmt.Println("Failed to convert MyFunction to func([]float64) float64")
		return
	}

	// Example data
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	// Apply the function to the data
	result := myFunction(data)
	fmt.Println(result)
}
