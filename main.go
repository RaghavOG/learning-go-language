package main

import "fmt"

func main() {
    // fmt.Println("Hello Raghav, Go is ready!")
	// var name string= "Raghav"
	// fmt.Println("Hello", name)

	// var age int = 21;
	// fmt.Println("Age:", age)

	// var height float64 = 1.75;
	// fmt.Println("Height:", height)

	// var isStudent bool = true;
	// fmt.Println("Is Student:", isStudent)
	// fmt.Println("Hello", name, "you are", age, "years old and", height, "meters tall. You are a student:", isStudent)


	// cgpa := 8;
	// fmt.Println("CGPA:", cgpa)


	// result := multiply(5, 6)
	// fmt.Println("Result:", result)

	// result2 := add(5, 6)
	// fmt.Println("Result2:", result2)


	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println( arr)
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }


	// nums := []int{10, 20, 30}
	// nums = append(nums, 40)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 50)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 60)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 70)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 80)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 90)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Length of nums:", len(nums))
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 90)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Length of nums:", len(nums))
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 90)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Length of nums:", len(nums))
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 90)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Length of nums:", len(nums))
	// fmt.Println("Capacity of nums:", cap(nums))
	// nums = append(nums, 90)
	// fmt.Println("Nums:", nums)
	// fmt.Println("Length of nums:", len(nums))
	// fmt.Println("Capacity of nums:", cap(nums))



	user := map[string]string{
		"name": "Raghav",
		"role": "Developer",
	}
	fmt.Println(user["name"]) // Raghav


	user["role"] = "Admin"
	fmt.Println(user["role"]) // Admin


}



func add(a int, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

