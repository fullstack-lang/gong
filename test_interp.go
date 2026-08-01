package main

import (
	"fmt"
	"math"
)

func main() {
	target := 22.5 * math.Pi / 180.0
	a1 := 0.0
	a2 := 90.0 * math.Pi / 180.0
	t := (target - a1) / (a2 - a1)
	x := 1.0 + t*(0.0 - 1.0)
	z := 0.0 + t*(1.0 - 0.0)
	angle := math.Atan2(z, x) * 180.0 / math.Pi
	fmt.Printf("t: %v, angle: %v\n", t, angle)
}
