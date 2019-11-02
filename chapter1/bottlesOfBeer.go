package chapter1

import "fmt"

func BottlesOfBeer(n int)  {
	for i := n;i > 0; i-=1 {
		fmt.Printf("%d bottles of beer on the wall, %d  bottles of beer\n", i,i)
		fmt.Printf("Take one down, pass it around, %d bottles of beer on the wall.\n", i-1)

	}

	fmt.Println("No bottles of beer on the wall, no bottles of beer,")
	fmt.Printf("Go to the store, buy some more, %d bottles of beer on the wall", n)
}
