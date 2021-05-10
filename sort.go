package main

import (
	"fmt"
	"time"
)
import "math/rand"

// main is the entry point for the program.
func main() {
	rand.Seed(time.Now().Unix())
	var arr []int
	arr = make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	fmt.Println(bubbles(arr))
	fmt.Println(bucket(arr))
	fmt.Println(quick(arr))
}

func quick(arr []int) []int {
	result := make([]int, len(arr))
	var tmp int
	copy(result, arr)
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	start := 0
	end := len(arr) - 1
	for ; start != end;
	{
		for ; arr[end] >= base && start != end; end-- {
		}
		for ; arr[start] <= base && start != end; start++ {
		}
		tmp = arr[start]
		arr[start] = arr[end]
		arr[end] = tmp
	}
	//交换初始位置与终点
	tmp = arr[0]
	arr[0] = arr[start]
	arr[start] = tmp
	//递归继续
	head := quick(arr[0:end])
	tail := quick(arr[end+1 : len(arr)])
	result = append(head, arr[end])
	result = append(result, tail...)
	return result
}

func bucket(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	index := 0
	var buckets [1000]int
	for i := 0; i < 1000; i++ {
		buckets[i] = 0
	}
	for i := 0; i < len(arr); i++ {
		buckets[arr[i]]++
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < buckets[i]; j++ {
			result[index] = i
			index++
		}
	}
	return result
}

func bubbles(arr []int) []int {
	var tmp int
	result := make([]int, len(arr))
	copy(result, arr)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			if result[i] < result[j] {
				tmp = result[i]
				result[i] = result[j]
				result[j] = tmp
			}
		}
	}
	return result
}
