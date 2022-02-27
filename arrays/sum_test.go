package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestSumAllTailsa(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	type args struct {
		numbersToSum [][]int
	}
	tests := []struct {
		name     string
		args     args
		wantSums []int
	}{
		{
			"make the sums of some slices",
			args{[][]int{[]int{1, 2}, []int{0, 9}}},
			[]int{2, 9},
		},
		{
			"safely sum empty slices",
			args{[][]int{[]int{}, []int{3, 4, 5}}},
			[]int{0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSums := SumAllTails(tt.args.numbersToSum...); !reflect.DeepEqual(gotSums, tt.wantSums) {
				t.Errorf("SumAllTails() = %v, want %v", gotSums, tt.wantSums)
			}
		})
	}
}
