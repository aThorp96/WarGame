package main

import (
	"fmt"
	cardDeck "github.com/aThorp96/WarGame/cardDeck"
)

type Game struct {
	player1 Hand
	player2 Hand
	deck    cardDeck.Deck
}

//Main function of the game.
func main() {
	player1 := *New()
	player2 := *New()
	deck := *cardDeck.NewDeck()
	game := Game{
		player1: player1,
		player2: player2,
		deck:    deck,
	}
	game.test()
	deck.Shuffle()

	for deck.HasNext() {
		card, _ := deck.Deal()
		player1.Give(card)
		card, _ = deck.Deal()
		player2.Give(card)
	}
}

func (game *Game) test() {
	game.deck.Sort()
	for i := 0; game.deck.HasNext(); i++ {
		card, _ := game.deck.Deal()
		fmt.Print(card.Rank() + card.Suit())
		if (1 != 0) && (i%13 == 0) {
			fmt.Println()
		}
	}
	game.deck.Shuffle()
	for i := 0; game.deck.HasNext(); i++ {
		card, _ := game.deck.Deal()
		fmt.Print(card.Rank() + card.Suit())
		if (1 != 0) && (i%13 == 0) {
			fmt.Println()
		}
	}
}
