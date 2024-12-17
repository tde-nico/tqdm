# tqdm

A Go version of Python tqdm library

## Functions

```go
func SetBarLength(length int) // set the length of the loading bar
func TqdmChan(start, end int64) <-chan int64 // returns a channel iterator
func Tqdm(arg1 int64, args ...int64) func(func(int64) bool) // returns a go iterator
```

## Examples

```go
for i := range tqdm.Tqdm(END) {}
for i := range tqdm.Tqdm(START, END) {}
for i := range tqdm.Tqdm(START, END, STEP) {}
```
