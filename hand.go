package main

import (
	"errors"
	cardDeck "github.com/aThorp96/WarGame/cardDeck"
	"math/rand"
)

type Hand struct {
	cards    [52]cardDeck.Card
	numCards int
}

//Creates, sorts, and returns a new hand.
func New() *Hand {
	hand := &Hand{}
	return hand
}

//Randomizes the order of the hand.
func (hand *Hand) Shuffle() {
	for i := 0; i < hand.numCards; i++ {
		index := rand.Intn(hand.numCards-i) + i
		temp := hand.cards[i]
		hand.cards[i] = hand.cards[index]
		hand.cards[index] = temp
	}
}

//Deal pops the top card off the hand and returns it.
//Returns an error if HasNext != true
func (hand *Hand) Deal() (*cardDeck.Card, error) {
	if !hand.HasNext() {
		return nil, errors.New("Hand is out of cards.")
	}
	hand.numCards--
	return &hand.cards[hand.numCards], nil
}

//hasNext returns true if there are cards left to be dealt.
func (hand *Hand) HasNext() bool {
	return hand.numCards > 0
}

//Give puts card into hand.
func (hand *Hand) Give(card *cardDeck.Card) {
	hand.cards[hand.numCards] = *card
	hand.numCards++
}
