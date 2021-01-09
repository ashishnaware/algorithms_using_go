package main

import "fmt"

func main()  {
	fmt.Println("Insertion Sort")

	input := [6]int{53, 12, 35, 21, 59, 15}
	fmt.Println("Input array: ", input)

	for i := 1; i < len(input);i++ {
		key := input[i]	// store current variable to key
		j := i-1		// pointing to previous record
		for j >= 0 && input[j] > key{
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
	fmt.Println("Output array: ", input)
}