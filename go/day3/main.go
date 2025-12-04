package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	sum2 := 0
	sum12 := 0

	for scn.Scan() {
		bank, err := parseBank(scn.Text(), 12)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//fmt.Printf("%s %d\n", scn.Text(), maxJoltagePickN(bank, 12))
		sum2 += maxJoltagePickN(bank, 2)
		sum12 += maxJoltagePickN(bank, 12)
	}

	fmt.Printf("pick 2: %d\npick 12: %d\n", sum2, sum12)
}

func maxJoltagePickN(bank []int, n int) int {
	if n < 1 {
		return 0
	}

	pow10 := []int{1, 1}
	for range n {
		pow10 = append(pow10, pow10[len(pow10)-1]*10)
	}

	memo := map[[2]int]int{}

	var rec func(int, int) int
	rec = func(i, digits int) int {
		if sum, ok := memo[[2]int{i, digits}]; ok {
			return sum
		}

		if digits == 0 {
			return 0
		}
		if i == len(bank) {
			return -pow10[len(pow10)-1] - 1
		}

		best := max(
			rec(i+1, digits),
			(pow10[digits]*bank[i])+rec(i+1, digits-1),
		)

		memo[[2]int{i, digits}] = best

		return best
	}

	return rec(0, n)
}

func parseBank(s string, size int) ([]int, error) {
	if len(s) < size {
		return nil, fmt.Errorf("bank %s is too short", s)
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return nil, fmt.Errorf("bank %s cannot contain rune %q", s, r)
		}
	}

	bank := []int{}
	for ns := range strings.SplitSeq(s, "") {
		n, _ := strconv.Atoi(ns)
		bank = append(bank, n)
	}

	return bank, nil
}
