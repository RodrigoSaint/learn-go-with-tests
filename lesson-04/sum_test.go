package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	expected := 15
	received := Sum(array)

	if expected != received {
		t.Errorf("Expected %d received %d", expected, received)
	}
}

func TestSumAll(t *testing.T) {
	received := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(received, want) {
		t.Errorf("Expected %v received %v", want, received)
	}
}
