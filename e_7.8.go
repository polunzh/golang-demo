package main

import "fmt"

func main()  {
	arr2 := []int{2, -1, 1, 3};
	fmt.Println(MinSlice(arr2))
	fmt.Println(MaxSlice(arr2))
}

func MinSlice(arr []int) (min int) {
	min = arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
	}

	return
}

func MaxSlice(arr []int) (min int) {
	min = arr[0]

	for _, value := range arr {
		if value > min {
			min = value
		}
	}

	return
}