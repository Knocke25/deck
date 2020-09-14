//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Suit is a set of all suits available in a normal deck of cards
type Suit uint8

const (
	Spade Suit = iota // iota does incrementing values. so Spade = 1, Diamond = 2, etc... heart = 4
	Diamond
	Club
	Heart
	Joker // Special
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank is a set of all ranks possible
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card is a set of Suit and Rank
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// Deck is a slice of cards
type Deck []Card

// New creates a new deck of 52 cards
func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}

	}
	// for each rank
	// add card{suit rank}
	return cards
}
