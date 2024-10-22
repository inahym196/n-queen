package nqueen

import (
	"fmt"
	"log/slog"
	"slices"
)

type nqueen struct{}

func SolveNqueen(n int) (solutions [][]int) {
	solutions = SolveNqueenRecursively(n, []int{})
	slog.Debug(fmt.Sprintf("total solutions: %v", solutions))
	return solutions
}

func SolveNqueenRecursively(n int, board []int) (solutions [][]int) {
	slog.Debug(fmt.Sprintf("call FindLocalNqueen(n:%d, board: %v)", n, board))
	if len(board) == n {
		slog.Debug(fmt.Sprintf("new solution: %v", board))
		return [][]int{board}
	}

	for x := 0; x < n; x++ {
		if slices.Contains(board, x) {
			continue
		}
		if onCrossLine(board, x) {
			continue
		}
		newBoard := make([]int, len(board), len(board)+1)
		copy(newBoard, board)
		newBoard = append(newBoard, x)
		newSolutions := SolveNqueenRecursively(n, newBoard)
		solutions = append(solutions, newSolutions...)
	}

	slog.Debug(fmt.Sprintf("return solutions: %v", solutions))
	return solutions
}

func onCrossLine(board []int, queen int) bool {
	queenRow := len(board)
	for i, q := range board {
		if AbsInt(queenRow-i) == AbsInt(queen-q) {
			return true
		}
	}
	return false
}

func AbsInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
