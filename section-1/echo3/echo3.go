package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Benchmark starts!")

	var start time.Time
	var s, sep string

	start = time.Now()
	for i := 0; i < 1000; i++ {
		s, sep = "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
	fmt.Printf("%v\nImplementation 2: %v ns elapsed\n", s, time.Since(start).Nanoseconds())

	start = time.Now()
	for i := 0; i < 1000; i++ {
		s, sep = "", ""
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
	fmt.Printf("%v\nImplementation 1: %v ns elapsed\n", s, time.Since(start).Nanoseconds())

	start = time.Now()
	for i := 0; i < 1000; i++ {
		s = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("%v\nImplementation 3: %v ns elapsed\n", s, time.Since(start).Nanoseconds())
}
