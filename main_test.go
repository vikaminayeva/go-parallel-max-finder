package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	size := 100
	result := generateRandomElements(size)

	if len(result) != size {
		t.Errorf("erorr %d, conclusion %d", size, len(result))
	}

	zeroResult := generateRandomElements(0)
	if zeroResult != nil {
		t.Errorf("error conclusion")
	}
}

func TestMaximum(t *testing.T) {
	testData := []int{10, 50, 20}
	expected := 50
	res := maximum(testData)
	if res != expected {
		t.Errorf("erorr %d, conclusion %d", expected, res)
	}
	if res := maximum([]int{}); res != 0 {
		t.Errorf("erorr %d", res)
	}
}
