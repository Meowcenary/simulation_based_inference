package util

import (
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

// basic functions for testing GetFunctionName
func NamedFunction(param int) int { return param }
func AnotherNamedFunction(param float64) float64 { return param }

func TestGetFunctionName(t *testing.T) {
	expectedName := "NamedFunction"
	name := GetFunctionName(NamedFunction)
	if expectedName != name {
		t.Errorf("Expected: %s\nGot: %s", expectedName, name)
	}

	expectedName = "AnotherNamedFunction"
	name = GetFunctionName(AnotherNamedFunction)
	if expectedName != name {
		t.Errorf("Expected: %s\nGot: %s", expectedName, name)
	}
}

func TestGenerateRandomVector(t *testing.T) {
	expectedVectorLength := 10
	generatedVector := GenerateRandomVector(random(), 10)
	generatedVectorLength := len(generatedVector)

	if expectedVectorLength != generatedVectorLength {
		t.Errorf("Expected vector of length: %d, Got vector of length: %d", expectedVectorLength, generatedVectorLength)
	}
}

func TestGenerateRandomMatrix(t *testing.T) {
	expectedRows := 3
	expectedColumns := 4
	generatedMatrix := GenerateRandomMatrix(random(), expectedRows, expectedColumns)
	generatedRows := len(generatedMatrix)
	generatedColumns := len(generatedMatrix[0])

	if expectedRows != generatedRows {
		t.Errorf("Expected %v rows\nGot %v rows", expectedRows, generatedRows)
	} else if expectedColumns != generatedColumns {
		t.Errorf("Expected %v rows\nGot %v rows", expectedColumns, generatedColumns)
	}
}
