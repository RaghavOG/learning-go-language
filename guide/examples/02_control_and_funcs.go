package main

import (
	"fmt"
	"strings"
)

// Run with: go run guide/examples/02_control_and_funcs.go
func main() {
	// if with short statement
	if n := 7; n%2 == 1 {
		fmt.Println("odd:", n)
	}

	// for: classic, while-style, range
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	j := 0
	for j < 3 {
		fmt.Print("j", j, " ")
		j++
	}
	fmt.Println()

	words := []string{"go", "is", "fun"}
	for idx, w := range words {
		fmt.Println(idx, w)
	}

	// switch (tagged)
	switch day := "mon"; day {
	case "mon", "tue":
		fmt.Println("weekday start")
	case "sat", "sun":
		fmt.Println("weekend")
	default:
		fmt.Println("midweek")
	}

	// switch (tagless)
	score := 82
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	defer fmt.Println("defer runs last in LIFO order")
	defer fmt.Println("defer 2")

	total, avg := sumAndAvg(10, 20, 30, 40)
	fmt.Println("total:", total, "avg:", avg)

	repeat := makeRepeater("go")
	fmt.Println(repeat(3))

	fmt.Println("multiply named result:", multiply(3, 4))
}

// Variadic with multiple returns
func sumAndAvg(nums ...int) (int, float64) {
	if len(nums) == 0 {
		return 0, 0
	}
	total := 0
	for _, n := range nums {
		total += n
	}
	return total, float64(total) / float64(len(nums))
}

// Returns a closure
func makeRepeater(word string) func(int) string {
	return func(n int) string {
		return strings.Repeat(word, n)
	}
}

// Named return values (use sparingly)
func multiply(a, b int) (product int) {
	product = a * b
	return // implicit return product
}

