package cardDeck


/*
* card stuct is a single card and has a suit and value.
* Suit is a number from 0:3. 0 = Hearts, 1 = Clubs, 2 = Diamonds, 3 = Spades.
* Rank is a number 2-14, with 11-13 being face cards and 14 being Ace. 
*/
type Card struct{
    suit int64
    rank int64
}

// Suit returns the suit of the card.
func (card *Card) Suit() int64 {
    return card.suit
}

//Rank returns the rank of the card.
func (card *Card) Rank() int64 {
    return card.rank
}

//New returns a new Card made from a passed in rank and suit (in that order).
func New(rank int64, suit int64) *Card {
    return &Card{suit, rank}
}



