package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Single Slice Sum", func(t *testing.T) {
		numbers := []int{3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 25

		if got != want {
			t.Errorf("got %d expected %d", got, want)
		}
	})

	t.Run("Multi Slice Sum", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{3, 4, 5, 6, 7}
		got := SumAll(slice1, slice2)
		want := []int{15, 25}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	}

	t.Run("sum of all tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("sum of all tails of slices, with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
