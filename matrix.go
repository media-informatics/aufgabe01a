package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	m = flag.Int("zeilen", 10, "Zeilen in der Matrix")
	n = flag.Int("spalten", 10, "Spalten in der Matrix")
)

func main() {
	flag.Parse()
	// compute largest needed column width for printing:
	format := fmt.Sprintf("%%%dd ", int(math.Log10(float64(*n**m))+1))

	// generate matrix with entries a[i,j] = i*j
	a := make([][]int, *m)
	for j := 0; j < *m; j++ {
		a[j] = make([]int, *n)
		for i := 0; i < *n; i++ {
			a[j][i] = i * j
		}
	}

	// print matrix to stdout
	for _, line := range a {
		for _, v := range line {
			fmt.Printf(format, v)
		}
		fmt.Println()
	}
}
