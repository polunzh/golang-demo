package main

import "fmt"

func main()  {
	arr1 := []float32{1.1, 2.1};
	arr2 := []int{1, 2};
	fmt.Println(Sum(arr1))
	fmt.Println(SumAndAverage(arr2))
}

func Sum(arr []float32) float32 {
	var res float32 = 0
	for _, value := range arr {
		res += value
	}

	return res
}

func SumAndAverage(arr []int) (sum int, average float32) {
	for _, value := range arr {
		sum += value
	}

	return sum, float32(sum / len(arr))
}