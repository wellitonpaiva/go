package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}

	got := Sum(numbers)
	expected := 21

	if got != expected {
		t.Errorf("Expected %d, but it was %v instead", expected, got)
	}
}

func TestSumWithZero(t *testing.T) {
	numbers := []int{0, 0, 0, 0, 0, 0}

	got := Sum(numbers)
	expected := 0

	if got != expected {
		t.Errorf("Expected %d, but it was %v instead", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
