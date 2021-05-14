package main

import "fmt"

type node struct {
	value   int
	pre     *node
	next    *node
	hasNext bool
}

func main() {
	var x []node
	base := []int{2, 3, 5, 8, 9, 10, 18, 26, 32}
	for i := 0; i < len(base); i++ {
		var tmp node
		tmp.value = base[i]
		x = append(x, tmp)
	}
	for i := 0; i < len(x)-1; i++ {
		x[i].next = &(x[i+1])
		x[i].hasNext = true
		if i != 0 {
			x[i].pre = &(x[i-1])
		}
	}
	insert(&x, 6)
	myNode := x[0]
	for {
		fmt.Println(myNode.value)
		if myNode.hasNext {
			myNode = *(myNode.next)
		} else {
			break
		}
	}
}

func insert(i *[]node, value int) {
	myNode := (*i)[0]
	for {
		if value < myNode.value {
			var tmp node
			tmp.value = value
			tmp.pre = myNode.pre
			tmp.next = &myNode
			tmp.hasNext = true
			myNode.pre.next = &tmp
			*i = append(*i, tmp)
			return
		}
		if myNode.hasNext {
			myNode = *(myNode.next)
		}
	}
}
