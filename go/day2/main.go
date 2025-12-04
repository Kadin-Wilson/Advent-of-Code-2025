package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(splitComma)

	twoSum := 0
	nSum := 0

	for scn.Scan() {
		first, last, err := parseRange(scn.Text())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, n := range repeatsTwice(first, last) {
			twoSum += n
		}

		for _, n := range repeatsN(first, last) {
			nSum += n
		}
	}

	fmt.Printf("two: %d\nn: %d\n", twoSum, nSum)
}

func repeatsTwice(first, last int) []int {
	nums := []int{}

	for i := first; i <= last; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 != 0 {
			continue
		}

		if s[:len(s)/2] == s[len(s)/2:] {
			nums = append(nums, i)
		}
	}

	return nums
}

func repeatsN(first, last int) []int {
	nums := []int{}

	for i := first; i <= last; i++ {
		s := strconv.Itoa(i)
		for j := 1; j <= len(s)/2; j++ {
			if repeats(s, j) && !slices.Contains(nums, i) {
				nums = append(nums, i)
			}
		}
	}

	//fmt.Printf("%d-%d %v\n", first, last, nums)

	return nums
}

func repeats(s string, n int) bool {
	if n < 1 {
		return false
	}
	if len(s)%n != 0 {
		return false
	}

	for i := n; i < len(s); i += n {
		if s[0:n] != s[i:i+n] {
			return false
		}
	}

	return true
}

func parseRange(s string) (int, int, error) {
	nums := strings.Split(s, "-")
	if len(nums) != 2 {
		return 0, 0, fmt.Errorf("failed to split range %s", s)
	}

	start, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%s in %s is not a number", nums[0], s)
	}

	end, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%s in %s is not a number", nums[1], s)
	}

	if start > end {
		return 0, 0, fmt.Errorf(
			"range start %d must be less than range end %d",
			start,
			end,
		)
	}

	return start, end, nil
}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		if len(data) > 0 {
			if data[len(data)-1] == '\n' {
				data = data[:len(data)-1]
			}
			return 0, data, bufio.ErrFinalToken
		}
		return 0, nil, bufio.ErrFinalToken
	}

	comma := slices.Index(data, ',')

	if comma == -1 {
		return 0, nil, nil
	}

	return comma + 1, data[:comma], nil
}
