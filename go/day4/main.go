package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	room, err := parseRoom(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	count := 0

	for rn, row := range room {
		for cn, col := range row {
			if col == '@' && adjacent(room, rn, cn) < 4 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func adjacent(room []string, row, col int) int {
	count := 0

	positions := []bool{
		row > 0 && room[row-1][col] == '@',
		col > 0 && room[row][col-1] == '@',
		row > 0 && col > 0 && room[row-1][col-1] == '@',
		row < len(room)-1 && room[row+1][col] == '@',
		col < len(room[0])-1 && room[row][col+1] == '@',
		row < len(room)-1 && col < len(room[0])-1 && room[row+1][col+1] == '@',
		row > 0 && col < len(room[0])-1 && room[row-1][col+1] == '@',
		col > 0 && row < len(room)-1 && room[row+1][col-1] == '@',
	}

	for _, b := range positions {
		if b {
			count++
		}
	}

	return count
}

func parseRoom(io.Reader) ([]string, error) {
	scn := bufio.NewScanner(os.Stdin)

	room := []string{}

	rowN := 0
	for scn.Scan() {
		rowN++
		row := scn.Text()
		if strings.ContainsFunc(row, invalidRoomRune) {
			return nil, fmt.Errorf("row #%d contains an invalid rune", rowN)
		}
		room = append(room, row)
	}

	if len(room) > 1 {
		l := len(room[1])
		for i, row := range room {
			if len(row) != l {
				return nil, fmt.Errorf("row #%d is too short", i+1)
			}
		}
	}

	return room, nil
}

func invalidRoomRune(r rune) bool { return r != '.' && r != '@' }
