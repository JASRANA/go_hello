package sum

import (
	"reflect"
	"testing"
)

// 覆盖率：-cover
func TestSum(t *testing.T) {

	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("any nums of collections of any size", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{2, 3, 4}

		got := SumAll(numbers1, numbers2)
		//want := "Bob" // 注意：DeepEqual 可以通过编译！
		want := []int{6, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v, given1 %v, given2 %v", got, want, numbers1, numbers2)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("make the sum of blank slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want :=[]int{0, 9}

		checkSums(t, got, want)
	})
}
