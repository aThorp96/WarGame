package cardDeck

import (
	"errors"
	"log"
	"math/rand"
)

type Deck struct {
	cards    [52]Card
	numCards int
}

func NewDeck() *Deck {
	deck := Deck{}
	deck.Shuffle()
	return &deck
}

//Creates a new deck in order.
func (deck *Deck) Sort() {
	for i := 0; i < 4; i++ {
		for j := 2; j < 15; j++ {
			deck.cards[(i+1)*j] = Card{i, j}
			deck.numCards++
		}
	}
}

//Randomizes the order of the deck.
func (deck *Deck) Shuffle() {
	deck.Sort()
	for i := 0; i < deck.numCards; i++ {
		index := rand.Intn(deck.numCards-i) + i
		temp := deck.cards[i]
		deck.cards[i] = deck.cards[index]
		deck.cards[index] = temp
	}
}

//Deal pops the top card off the deck and returns it.
//Returns an error if HasNext != true
func (deck *Deck) Deal() (*Card, error) {
	if !deck.HasNext() {
		return nil, errors.New("Deck is out of cards.")
	}
	deck.numCards--
	return &deck.cards[deck.numCards], nil
}

//hasNext returns true if there are cards left to be dealt.
func (deck *Deck) HasNext() bool {
	return deck.numCards > 0
}
