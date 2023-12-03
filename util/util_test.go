package util

import (
	"testing"
)

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
