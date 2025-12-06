package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)

	r, err := scanRanges(scn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ids, err := scanIDs(scn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fresh := 0
	for _, id := range ids {
		if r.inRange(id) {
			fresh++
		}
	}

	fmt.Println(fresh)
}

type ranges struct {
	count int
	low   []int
	high  []int
}

func (r *ranges) inRange(n int) bool {
	for i, l := range r.low {
		h := r.high[i]
		if n >= l && n <= h {
			return true
		}
	}
	return false
}

func parseRange(s string) (int, int, error) {
	nums := strings.Split(s, "-")
	if len(nums) != 2 {
		return 0, 0, fmt.Errorf("range %s is improperly formatted", s)
	}

	l, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%s in %s is not a number", nums[0], s)
	}
	r, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%s in %s is not a number", nums[1], s)
	}

	if l > r {
		return 0, 0, fmt.Errorf("low val %d is greater than high val %d", l, r)
	}

	return l, r, nil
}

func scanRanges(scn *bufio.Scanner) (*ranges, error) {
	r := ranges{
		low:  []int{},
		high: []int{},
	}

	for scn.Scan() {
		if len(scn.Text()) == 0 {
			break
		}

		r.count++

		l, h, err := parseRange(scn.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse range #%d: %w", r.count, err)
		}

		r.low = append(r.low, l)
		r.high = append(r.high, h)
	}

	return &r, nil
}

func scanIDs(scn *bufio.Scanner) ([]int, error) {
	ids := []int{}

	for scn.Scan() {
		id, err := strconv.Atoi(scn.Text())
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse id #%d, %s",
				len(ids)+1,
				scn.Text(),
			)
		}

		ids = append(ids, id)
	}

	return ids, nil
}
