package main

import (
	"fmt"
	 . "github.com/HJPAY/holdem/cards"
)

func main() {
	fmt.Println("holdem!!")
	var deck *Deck = NewDeck()

	for _, card := range deck.Cards {
		fmt.Println(card.String())
	}
}
