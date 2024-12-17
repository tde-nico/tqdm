package tqdm

import (
	"fmt"
)

var barLength = 100

func SetBarLength(length int) {
	barLength = length
}

func TqdmChan(start, end int64) <-chan int64 {
	ch := make(chan int64)
	total := end - start
	pb := NewProgressBar(total)

	go func() {
		defer close(ch)
		ch <- start
		pb.Start()
		for i := start + 1; i < end; i++ {
			ch <- i
			pb.Update()
		}
		pb.Update()
		fmt.Println()
	}()

	return ch
}

func Tqdm(arg1 int64, args ...int64) func(func(int64) bool) {
	var start, end, step int64
	if len(args) == 0 {
		start = 0
		end = arg1
		step = 1
	} else if len(args) == 1 {
		start = arg1
		end = args[0]
		step = 1
	} else if len(args) == 2 {
		start = arg1
		end = args[0]
		step = args[1]
	} else {
		panic("too many arguments")
	}

	pb := NewProgressBar((end - start) / step)

	return func(yield func(int64) bool) {
		pb.Start()
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
			pb.Update()
		}
		fmt.Println()
	}
}
