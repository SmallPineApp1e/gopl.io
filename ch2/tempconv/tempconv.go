// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// practice2.2
type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func (f Foot) String() string     { return fmt.Sprintf("%g 英尺", f) }
func (m Meter) String() string    { return fmt.Sprintf("%g 米", m) }
func (p Pound) String() string    { return fmt.Sprintf("%g 英镑", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g 公斤", k) }

//!-
