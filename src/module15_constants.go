package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Printf("Hello type %T\n", World)
	fmt.Printf("Happy type %T\n", Pi, "Day")

	const Truth = true
	fmt.Printf("Go rules? type %T\n", Truth)
}
