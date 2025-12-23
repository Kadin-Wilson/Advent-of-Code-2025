package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type operator func(...int) int

var operators map[rune]operator = map[rune]operator{
	'+': multiAdd,
	'*': multiMul,
}

func main() {
	problems, err := parseProblems(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
	}

	total := 0
	for _, p := range problems {
		total += p.solve()
	}

	fmt.Println(total)
}

type problem struct {
	operands  []int
	operation operator
}

func (p problem) solve() int {
	return p.operation(p.operands...)
}

func parseProblems(in io.Reader) ([]problem, error) {
	problems := []problem{}

	scn := bufio.NewScanner(in)
	for scn.Scan() {
		tokens := strings.Fields(scn.Text())

		for len(tokens) > len(problems) {
			problems = append(problems, problem{operands: []int{}})
		}

		if isOperandLine(scn.Text()) {
			for i, tok := range tokens {
				n, err := strconv.Atoi(tok)
				if err != nil {
					return nil, err
				}
				problems[i].operands = append(problems[i].operands, n)
			}
		} else if isOperatorLine(scn.Text()) {
			for i, tok := range tokens {
				if len(tok) != 1 {
					return nil, fmt.Errorf("operator %s is invalid", tok)
				}
				problems[i].operation = operators[[]rune(tok)[0]]
			}
		} else {
			return nil, fmt.Errorf("invalid problem line")
		}
	}

	return problems, nil
}

func isOperandLine(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsSpace(r) && !unicode.IsNumber(r)
	})
}

func isOperatorLine(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		_, ok := operators[r]
		return !ok && !unicode.IsSpace(r)
	})
}

func multiAdd(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func multiMul(nums ...int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}
