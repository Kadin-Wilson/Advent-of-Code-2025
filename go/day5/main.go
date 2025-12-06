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

	fmt.Printf(
		"available: %d\npossible: %d\n",
		r.availableCount(ids),
		r.possibleCount(),
	)
}

type ranges struct {
	low  []int
	high []int
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

func (r *ranges) availableCount(ids []int) int {
	count := 0
	for _, id := range ids {
		if r.inRange(id) {
			count++
		}
	}
	return count
}

func (r *ranges) possibleCount() int {
	count := 0
	for i, l := range r.low {
		h := r.high[i]
		count += h - l + 1
	}
	return count
}

type pair struct {
	low  int
	high int
}

func (p pair) overlap(pp pair) bool {
	if p.low >= pp.low && p.low <= pp.high {
		return true
	}
	if p.high >= pp.low && p.high <= pp.high {
		return true
	}
	if p.low <= pp.low && p.high >= pp.high {
		return true
	}
	return false
}

func (p pair) join(pp pair) (pair, error) {
	if !p.overlap(pp) {
		return pair{}, fmt.Errorf("pairs %v and %v do not overlap", p, pp)
	}
	return pair{low: min(p.low, pp.low), high: max(p.high, pp.high)}, nil
}

func parsePair(s string) (pair, error) {
	nums := strings.Split(s, "-")
	if len(nums) != 2 {
		return pair{}, fmt.Errorf("range %s is improperly formatted", s)
	}

	l, err := strconv.Atoi(nums[0])
	if err != nil {
		return pair{}, fmt.Errorf("%s in %s is not a number", nums[0], s)
	}
	h, err := strconv.Atoi(nums[1])
	if err != nil {
		return pair{}, fmt.Errorf("%s in %s is not a number", nums[1], s)
	}

	if l > h {
		return pair{}, fmt.Errorf("low val %d is greater than high val %d", l, h)
	}

	return pair{low: l, high: h}, nil
}

func scanRanges(scn *bufio.Scanner) (*ranges, error) {
	count := 0

	rm := map[pair]struct{}{}

	for scn.Scan() {
		if len(scn.Text()) == 0 {
			break
		}

		count++

		p, err := parsePair(scn.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse range #%d: %w", count, err)
		}

		for pp := range rm {
			if p.overlap(pp) {
				p, _ = p.join(pp)
				delete(rm, pp)
			}
		}

		rm[p] = struct{}{}
	}

	r := ranges{
		low:  []int{},
		high: []int{},
	}

	for p := range rm {
		r.low = append(r.low, p.low)
		r.high = append(r.high, p.high)
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
