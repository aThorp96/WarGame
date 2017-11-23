package cardGame

import "math/rand"

type Deck struct {
    cards [52]Card
    numCards int64
}

// NoCard will be returned if a card is attempted to be accessed, but there is no card.
type NoCard struct {
    Error string
}

//Creates a new deck in order. 
func (deck *Deck) Sort() {
    for i := 0; i < 4; i++ {
        for j := 2; j < 15; j++ {
            deck.cards[(i + 1) * j] = card.New(i, j)
            deck.numCards++
        }
    }
}

//Creates, sorts, and returns a new deck. 
func New() *Deck {
    deck := &Deck{Card[52], 0}
    deck.Sort()
    return deck
}

//Randomizes the order of the deck. 
func (deck *Deck) Shuffle() {
    deck.Sort()
    for i := 0; i < deck.numCards; i++ {
        index := rand.Intn(deck.numCards - i) + i
        temp := deck.cards[i]
        deck.cards[i] = deck.cards[index]
        deck.cards[index] = temp
    }
}

//Deal pops the top card off the deck and returns it.
//Returns an error if HasNext != true
func (deck *Deck) Deal() (*Card, error) {
    if !deck.HasNext() {
        return nil, NoCard{"Deck is out of cards."}
    }
    deck.numCards--
    return deck.cards[numCards], nil
}

//hasNext returns true if there are cards left to be dealt.
func (deck *Deck) HasNext() bool {
    return numCards > 0
}

