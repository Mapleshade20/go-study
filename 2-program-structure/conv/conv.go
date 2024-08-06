package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Celsius    float64
	Fahrenheit float64
	Feet       float64
	Meter      float64
	Pound      float64
	Kilogram   float64
)

var (
	t = flag.Bool("t", false, "temperature")
	l = flag.Bool("l", false, "length")
	w = flag.Bool("w", false, "weight")
)

func main() {
	flag.Parse()
	strNums := flag.Args()

	if len(strNums) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter space-separated numbers: ")
		scanner.Scan()
		input := scanner.Text()
		strNums = strings.Fields(input)
	}
	convert(strNums)
}

func convert(strNums []string) {
	for _, strNum := range strNums {
		num, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			fmt.Printf("Error parsing '%s': %v", strNum, err)
			continue
		}
		switch {
		case *t:
			fmt.Printf("%v F = %.1f C, %v C = %.1f F", num, FToC(Fahrenheit(num)), num, CToF(Celsius(num)))
		// do likewise
		default:
			fmt.Print("Error: no flags set")
		}
	}
}

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
