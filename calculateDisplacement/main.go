package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a0, v0, s0 float64) func(t float64) float64 {
	fn := func(time float64) float64 {
		return (0.5*(math.Pow(time, 2)*a0) + (time * v0) + s0)
	}
	return fn
}

func main() {
	// array := []int{1, 3, 5, 4, 2}

	var acc float64
	fmt.Println("Enter acceleration:")
	fmt.Scanln(&acc)

	var init_vel float64
	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&init_vel)

	var init_disp float64
	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&init_disp)

	var time float64
	fmt.Println("Enter time:")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acc, init_vel, init_disp)

	ans := fn(time)
	fmt.Printf("Answer is %.2f\n", ans)

}
