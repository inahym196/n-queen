package nqueen_test

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	nqueen "github.com/inahym196/n-queen"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	})))
}

func TestForQueen(t *testing.T) {
	want := [][]int{
		{1, 3, 0, 2},
		{2, 0, 3, 1},
	}
	sols := nqueen.SolveNqueen(4)

	if len(sols) != 2 {
		t.Errorf("num of solution should be 2, but %d", len(sols))
	}
	if !reflect.DeepEqual(want, sols) {
		t.Errorf("want %v, but %v", want, sols)
	}

}
