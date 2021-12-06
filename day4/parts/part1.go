package parts

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_for_win(board Board) bool {
	for _, row := range board {
		xs := 0
		for _, num := range row {
			if num[0] == 'x' {
				xs += 1
			}
		}
		if xs == len(row) {
			return true
		}
	}

	for i := range board[0] {
		xs := 0
		for _, row := range board {
			if row[i][0] == 'x' {
				xs += 1
			}
		}
		if xs == len(board) {
			return true
		}
	}

	return false
}

func calc_score(board Board, num_string string) int {
	unmarked_score := 0
	num, err := strconv.Atoi(num_string)
	check(err)
	for _, row := range board {
		for _, cell := range row {
			if cell[0] != 'x' {
				val, err := strconv.Atoi(cell)
				check(err)
				unmarked_score += val
			}
		}
	}
	return unmarked_score * num
}

type Board = [][]string

func board_from_string(boardS string) Board {
	var string_rows = strings.Split(boardS, "\n")
	var board [][]string
	for _, row := range string_rows {
		var new_row []string
		for _, num := range strings.Split(row, " ") {
			clean_num := strings.TrimSpace(num)
			if len(clean_num) > 0 {
				new_row = append(new_row, clean_num)
			}
		}
		board = append(board, new_row)
	}
	return board
}

func mark_board(board Board, num string) Board {
	var new_board Board

	for _, row := range board {
		var new_row []string
		for _, board_num := range row {
			if board_num == num {
				new_row = append(new_row, "x"+board_num)
			} else {
				new_row = append(new_row, board_num)
			}
		}
		new_board = append(new_board, new_row)
	}
	return new_board
}

func Part1(input_file string) int {
	fmt.Println("Running Part1:", input_file)
	data, err := ioutil.ReadFile(input_file)
	check(err)
	blocks := strings.Split(string(data), "\n\r")
	var bingo_data []string
	for _, line := range blocks {
		bingo_data = append(bingo_data, strings.TrimSpace(line))
	}

	numbers := strings.Split(bingo_data[0], ",")

	var boards []Board
	for _, boardS := range bingo_data[1:] {
		boards = append(boards, board_from_string(boardS))
	}

	var score int
	for _, number := range numbers {
		var tmp_boards []Board
		for _, board := range boards {
			new_board := mark_board(board, number)
			if check_for_win(new_board) {
				fmt.Println(new_board)
				score = calc_score(new_board, number)
				break
			}
			tmp_boards = append(tmp_boards, new_board)
		}
		if score > 0 {
			fmt.Println(score)
			break
		}
		boards = tmp_boards
	}
	return score
}
