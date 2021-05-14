package main

import "fmt"

func main() {
	s1 := []int{2, 4, 1, 2, 5, 6}
	s2 := []int{3, 1, 3, 5, 6, 4}
	table := []int{}
	for {
		if len(s1) > 0 && len(s2) > 0 {
			putToTable(&table, &s1)
			putToTable(&table, &s2)
		} else {
			fmt.Println("===================")
			fmt.Println("table:", table)
			fmt.Println(s1)
			fmt.Println(s2)
			break
		}
		fmt.Println("table:", table)
		fmt.Println(s1)
		fmt.Println(s2)
	}
}

func putToTable(table *[]int, s *[]int) {
	//出牌
	currentCard := (*s)[0]

	//检测
	for i := 0; i < len(*table); i++ {
		if (*table)[i] == currentCard {
			*s = append((*s)[:0], (*s)[1:]...)
			*table = append(*table, currentCard)
			reversed := reverse((*table)[i:])
			//reversed := (*table)[i:]
			*s = append(*s, reversed...)
			*table = (*table)[:i]
			return
		}
	}
	*s = append((*s)[:0], (*s)[1:]...)
	*table = append(*table, currentCard)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
