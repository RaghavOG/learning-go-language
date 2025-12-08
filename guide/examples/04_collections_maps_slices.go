package main

import "fmt"

// Run with: go run guide/examples/04_collections_maps_slices.go
func main() {
	// Slices: length, capacity, append growth
	nums := []int{1, 2, 3}
	fmt.Println("nums:", nums, "len:", len(nums), "cap:", cap(nums))
	nums = append(nums, 4, 5)
	fmt.Println("after append:", nums, "len:", len(nums), "cap:", cap(nums))

	// Sub-slicing shares backing array
	part := nums[1:3]
	part[0] = 99 // affects original
	fmt.Println("part:", part, "nums now:", nums)

	// Copy to avoid sharing
	copyBuf := make([]int, len(part))
	copy(copyBuf, part)
	copyBuf[0] = 7
	fmt.Println("copied slice:", copyBuf, "original nums:", nums)

	// Map operations
	users := map[string]int{"alice": 1}
	users["bob"] = 2
	if id, ok := users["carol"]; !ok {
		fmt.Println("carol missing, ok =", ok, "id =", id)
	}
	delete(users, "alice")
	for name, id := range users {
		fmt.Println("user", name, "id", id)
	}

	// Array literal and pointer to array
	arr := [...]string{"go", "rust", "python"}
	var pArr *[3]string = &arr
	fmt.Println("array:", arr, "pointer sees:", *pArr)
}

