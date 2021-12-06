package parts

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2(input_file string) int {
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
		var new_board Board
		for _, board := range boards {
			new_board = mark_board(board, number)
			if !check_for_win(new_board) {
				tmp_boards = append(tmp_boards, new_board)
			}
		}
		if len(tmp_boards) == 0 {
			score = calc_score(new_board, number)
			fmt.Println(score, number)
			fmt.Println(new_board)
			break
		}
		boards = tmp_boards
	}
	return score
}
