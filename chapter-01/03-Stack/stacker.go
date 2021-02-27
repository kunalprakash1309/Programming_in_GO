package main

import (
	"Programming_in_GO/chapter-01/03-stack/stack"
	"fmt"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)
	fmt.Println(haystack)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}