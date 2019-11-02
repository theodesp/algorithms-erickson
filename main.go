package main

import (
	"fmt"
	"github.com/theodesp/algorithms-erickson/chapter1"
)

func main()  {
	votes := []int{100000, 80000, 30000}
	seats := 8
	fmt.Println(chapter1.EqualProportions(votes, seats))
}
