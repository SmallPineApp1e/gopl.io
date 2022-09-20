// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

const (
	KILO_TYPE  string = "k"
	POUND_TYPE string = "p"
	FOOT_TYPE  string = "f"
	METER_TYPE string = "m"
)

func main() {
	// example()
	practice2()
}

func example() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), c, tempconv.CToK(c))
	}
}

func practice2() {
	length := len(os.Args)
	fmt.Printf("os args length: %d\n", length)
	for i := 1; i < length-1; i += 2 {
		arg_type := os.Args[i]
		arg, err := strconv.ParseFloat(os.Args[i+1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		if arg_type == KILO_TYPE {
			kilo := tempconv.Kilogram(arg)
			fmt.Printf("%s = %s\n", kilo, tempconv.Kilo2Pound(kilo))
		} else if arg_type == POUND_TYPE {
			pound := tempconv.Pound(arg)
			fmt.Printf("%s = %s\n", pound, tempconv.Pound2Kilo(pound))
		} else if arg_type == FOOT_TYPE {
			foot := tempconv.Foot(arg)
			fmt.Printf("%s = %s\n", foot, tempconv.Foot2Meter(foot))
		} else if arg_type == METER_TYPE {
			meter := tempconv.Meter(arg)
			fmt.Printf("%s = %s\n", meter, tempconv.Meter2Foot(meter))
		}

	}
}

//!-
