package main

import (
	"fmt"
	"unsafe"
)

// Run with: go run guide/examples/01_vars_and_types.go
func main() {
	// Full form
	var name string = "Raghav"
	var age int = 21
	var height float64 = 1.75
	var isStudent bool = true

	// Short form (inside functions only)
	cgpa := 8 // type is inferred

	// Zero values (when only declared)
	var zeroInt int        // 0
	var zeroString string  // ""
	var zeroBool bool      // false
	var zeroSlice []int    // nil
	var zeroMap map[int]int // nil

	// Constants
	const country string = "India"
	const phi = 1.618 // untyped constant
	const (
		StatusOK    = 200
		StatusError = 500
	)

	// Numeric types and conversion
	var small uint8 = 255
	var big uint32 = uint32(small) // explicit conversion only

	// Rune (Unicode code point) and byte (alias for uint8)
	var letter rune = 'G'
	var b byte = 'A'

	fmt.Println("name:", name, "age:", age, "height:", height, "student:", isStudent)
	fmt.Println("cgpa:", cgpa, "country:", country, "phi:", phi)
	fmt.Println("zero values:", zeroInt, zeroString, zeroBool, zeroSlice, zeroMap)
	fmt.Println("converted:", small, "->", big)
	fmt.Println("rune:", letter, "byte:", b)

	// Arrays (fixed size) vs slices (dynamic view)
	var fixed [3]int = [3]int{1, 2, 3}
	flexible := []int{10, 20, 30}
	flexible = append(flexible, 40)
	fmt.Println("array:", fixed)
	fmt.Println("slice:", flexible, "len:", len(flexible), "cap:", cap(flexible))

	// make for slices and maps
	grades := make([]int, 0, 5)
	grades = append(grades, 9, 8, 10)
	fmt.Println("grades:", grades, "len:", len(grades), "cap:", cap(grades))

	user := make(map[string]string)
	user["name"] = "Raghav"
	user["role"] = "Developer"
	fmt.Println("user map:", user, "role lookup:", user["role"])

	// Strings are immutable; iterating by rune vs byte
	word := "Go→世界"
	fmt.Println("bytes:", []byte(word))
	for i, r := range word {
		fmt.Printf("rune %d: %c\n", i, r)
	}

	// Size of types (implementation dependent)
	fmt.Println("size of int (bytes):", unsafe.Sizeof(int(0)))
}

