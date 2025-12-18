package main

import "test/counter"

func main() {
	counter.Interval = 2
	counter.Count()
	counter.Count()
	counter.Count()
	counter.Count()
	counter.Count()
	counter.Print()
}
