package main

import "fmt"

func swap(x, y string, z int) (string, string, int) {
	return y, x, z
}

func main() {
	a, b, c := swap("hello", "world", 20)
	fmt.Println(a, b, c)
}
