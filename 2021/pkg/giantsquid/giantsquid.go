package giantsquid

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (p *Puzzle) PartOne() string {
	boards := convertBoards()
	for _, drawnNumber := range drawnNumbers {
		for _, board := range boards {
			board = handleDrawnNumber(drawnNumber, board)
			if checkBoard(board) {
				score := scoreBoard(board)
				return fmt.Sprintf("%v", score*drawnNumber)
			}
		}
	}
	return ""
}

func (p *Puzzle) PartTwo() string {
	boards := convertBoards()
	boardFlag := map[int]bool{}
	winnerCount := 0
	for _, drawnNumber := range drawnNumbers {
		for b, board := range boards {
			board = handleDrawnNumber(drawnNumber, board)
			if checkBoard(board) {
				if !boardFlag[b] {
					boardFlag[b] = true
					winnerCount++
				}
				if winnerCount == len(boards)-1 {
					score := scoreBoard(board)
					return fmt.Sprintf("%v", score*drawnNumber)
				}
			}
		}
	}
	return ""
}

func convertBoards() [][][]int {
	result := [][][]int{}

	for _, board := range inputs {
		rows := strings.Split(board, "\n")
		boardBuffer := [][]int{}
		for _, row := range rows {

			space := regexp.MustCompile(`\s+`)
			row = space.ReplaceAllString(row, " ")
			columns := strings.Split(row, " ")
			intColumns := []int{}
			for _, column := range columns {
				v, _ := strconv.Atoi(column)
				intColumns = append(intColumns, v)
			}

			boardBuffer = append(boardBuffer, intColumns)
		}

		result = append(result, boardBuffer)
	}
	return result
}

func handleDrawnNumber(dn int, board [][]int) [][]int {
	for r, rows := range board {
		for c, _ := range rows {
			if board[r][c] == dn {
				board[r][c] = -1
			}
		}
	}
	return board
}

func checkBoard(board [][]int) bool {
	// row check
	for r := 0; r < len(board); r++ {
		rowSum := 0
		for c := 0; c < len(board[0]); c++ {
			rowSum = rowSum + board[r][c]
		}
		if rowSum == -5 {
			return true
		}
	}

	for c := 0; c < len(board[0]); c++ {
		columnSum := 0
		for r := 0; r < len(board); r++ {
			columnSum = columnSum + board[r][c]
		}
		if columnSum == -5 {
			return true
		}
	}

	return false
}

func scoreBoard(board [][]int) int {
	score := 0
	for r, rows := range board {
		for c := range rows {
			if board[r][c] == -1 {
				continue
			}
			score = score + board[r][c]
		}
	}

	return score
}
