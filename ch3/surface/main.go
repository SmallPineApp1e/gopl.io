// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	width, height = 1280, 720           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func drawSurface(out http.ResponseWriter) {
	resource_str := ""
	resource_str += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			result, ax, ay := corner(i+1, j)
			if result != true {
				continue
			}
			result, bx, by := corner(i, j)
			if result != true {
				continue
			}
			result, cx, cy := corner(i, j+1)
			if result != true {
				continue
			}
			result, dx, dy := corner(i+1, j+1)
			if result != true {
				continue
			}
			resource_str += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	resource_str += fmt.Sprintf("</svg>")
	out.Header().Set("Content-Type", "image/svg+xml")
	out.Write([]byte(resource_str))
}

func corner(i, j int) (bool, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) {
		return false, -1, -1
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return true, sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	// return water(x, y, r)
	return egg(x, y, r)
}

func water(x, y, r float64) float64 {
	return math.Sin(r) / r
}

func egg(x, y, r float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
}

//!-
