package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	//cards.shuffle()
	//cards.print()

	a, b := deal(cards, 5)

	fmt.Println(cards)
	fmt.Println(a)
	fmt.Println(b)
}
