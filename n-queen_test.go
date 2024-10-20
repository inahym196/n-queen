package nqueen_test

import (
	"reflect"
	"testing"

	nqueen "github.com/inahym196/n-queen"
)

func TestForQueen(t *testing.T) {
	want := [][]int{
		{1, 3, 0, 2},
		{2, 0, 3, 1},
	}
	sols, err := nqueen.SolveNqueen(4)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(sols) != 2 {
		t.Errorf("num of solution should be 2, but %d", len(sols))
	}
	if !reflect.DeepEqual(want, sols) {
		t.Errorf("want %v, but %v", want, sols)
	}

}
