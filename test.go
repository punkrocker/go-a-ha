package main

import "fmt"

func main() {
	arr := []int{6, 3, 1, 7, 5, 8, 9, 2, 4}
	arr = append(arr[2:], arr[1])
	fmt.Println(arr)
}
