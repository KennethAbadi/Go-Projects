//go:generate stringer -type=Suit,Value
package deck

type Suit uint8 //Spade Diamond Club Heart, Joker for extra

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Value uint8 //2,3,4,5,6,7,8,9,10,j,q,k,a

const (
	x Value = iota //blank value
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
	Value
}

func (c Card) String() string {

}
