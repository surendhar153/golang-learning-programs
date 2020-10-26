package main

import "fmt"

func main() {
	var radius, volume float64
	const pi float64 = 3.14
	fmt.Println("Find Volume of Sphere of radius R")
	fmt.Scanln(&radius)
	volume = (float64(4) / float64(3)) * pi * radius * radius * radius
	fmt.Printf("%.2f", volume)
}
