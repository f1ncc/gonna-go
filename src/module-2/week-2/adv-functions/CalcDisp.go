package main

import (
	"fmt"
)

func GenDisplaceFn (a,initVel,initDisp float64) func(float64) float64 {
	fn := func (t float64) float64 {
		return ((a*t*t)/2 + initVel*t + initDisp)
	}
	return fn
} 

func main() {
	var (
		acc, initialvel, initialDisp, time float64
	)
	fmt.Printf("enter acceleration\n")
	fmt.Scan(&acc)

	fmt.Printf("enter initial velocity\n")
	fmt.Scan(&initialvel)

	fmt.Printf("enter initial displacement\n")
	fmt.Scan(&initialDisp)

	fmt.Printf("enter total time\n")
	fmt.Scan(&time)

	calcDisp := GenDisplaceFn(acc, initialvel, initialDisp)

	fmt.Printf("Final Displacement equals %f mts\n",calcDisp(time))
}