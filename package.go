package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
	sayHello()
	a, b := swap("站三", "李四")

	fmt.Println("a=", a, "b=", b)

	fmt.Println(8)
}

func sayHello() {
	var res = add(1, 22)
	fmt.Println("hello...", math.Pi, res)

}
func add(a int, b int) int {
	return a + b
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(s string) (s1, s2 string) {
	parts := strings.Split(s, ",")
	s1 = parts[0]
	s2 = parts[1]
	return
}
