package main

import "fmt"

func main() {
	arr := []int{6, 3, 1, 7, 5, 8, 9, 2, 4}
	fmt.Println(change(arr))
}

func change(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	index := 0
	for ; len(arr) != 0; {
		result[index] = arr[0]
		index++
		if len(arr) > 1 {
			arr = append(arr[2:], arr[1])
		} else {
			break
		}
		fmt.Println(arr)
	}

	return result
}
