//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota // iota does incrementing values. so Spade = 1, Diamond = 2, etc... heart = 4
	Diamond
	Club
	Heart
	Joker
)

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
