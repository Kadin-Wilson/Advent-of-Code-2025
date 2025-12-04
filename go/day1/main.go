package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	mp := newMultiProcessor(new(finalZeros), new(allZeros))
	counts, err := processInstructionSequence(os.Stdin, newDial(), &mp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("final zeros: %d\nall zeros: %d\n", counts[0], counts[1])
}

type dial int

func newDial() dial {
	return 50
}

func (d dial) rotate(offset int) dial {
	if offset >= 0 {
		return (d + dial(offset)) % 100
	}
	return (d + 100 + dial(offset)) % 100
}

func offsetFromString(s string) (int, error) {
	i, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, fmt.Errorf("value %s is not a number", s[1:])
	}
	if i < 0 {
		return 0, fmt.Errorf("value %d must be positive", i)
	}

	switch s[0] {
	case 'L':
		return -i, nil
	case 'R':
		return i, nil
	}

	return 0, errors.New("instruction must start with L or R")
}

type processor[T any] interface {
	Process(dial, int) dial
	Done() T
}

func processInstructionSequence[T any](
	in io.Reader,
	d dial,
	p processor[T],
) (T, error) {
	scn := bufio.NewScanner(in)
	lines := 0

	for scn.Scan() {
		lines++

		offset, err := offsetFromString(scn.Text())
		if err != nil {
			return *new(T), fmt.Errorf(
				"couldn't process instruction %d: %w",
				lines,
				err,
			)
		}

		d = p.Process(d, offset)
	}
	return p.Done(), nil
}

type finalZeros int

func (fz *finalZeros) Process(d dial, offset int) dial {
	d = d.rotate(offset)
	if d == 0 {
		*fz++
	}
	return d
}

func (fz *finalZeros) Done() int {
	return int(*fz)
}

type allZeros int

func (az *allZeros) Process(d dial, offset int) dial {
	rotations := offset / 100
	if rotations < 0 {
		rotations = -rotations
	}
	*az += allZeros(rotations)

	if d == 0 {
		return d.rotate(offset)
	}

	old := d
	d = d.rotate(offset)

	if d == 0 {
		*az++
	} else if offset < 0 && d > old {
		*az++
	} else if offset > 0 && d < old {
		*az++
	}

	return d
}

func (az *allZeros) Done() int {
	return int(*az)
}

type multiprocessor[T any] struct {
	procs []processor[T]
	dials []dial
}

func newMultiProcessor[T any](procs ...processor[T]) multiprocessor[T] {
	var mp multiprocessor[T]
	mp.procs = procs

	mp.dials = []dial{}
	for range len(procs) {
		mp.dials = append(mp.dials, newDial())
	}

	return mp
}

func (mp *multiprocessor[T]) Process(d dial, offset int) dial {
	for i, p := range mp.procs {
		mp.dials[i] = p.Process(mp.dials[i], offset)
	}
	return d.rotate(offset)
}

func (mp *multiprocessor[T]) Done() []T {
	rtr := []T{}
	for _, p := range mp.procs {
		rtr = append(rtr, p.Done())
	}
	return rtr
}
