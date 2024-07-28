package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("should work on a single slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumAll(numbers)
		want := []int{6}

		checkSums(t, got, want)
	})
	t.Run("should accept arbitary number of slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		numbers2 := []int{4, 5, 6}

		got := SumAll(numbers, numbers2)
		want := []int{6, 15}

		checkSums(t, got, want)
	})

}
