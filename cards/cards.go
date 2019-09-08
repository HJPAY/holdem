package cards

type Suits int

const (
	_ Suits = iota
	Spade
	Heart
	Diamond
	Club
)

func (s Suits) String() string {
	switch s {
	case Spade:
		return "s" //"Spade"
	case Heart:
		return "h" //"Heart"
	case Diamond:
		return "d" //"Diamond"
	case Club:
		return "c" //"Club"
	default:
		return "Unknown"
	}
}

type Number int

func (n Number) String() string {
	switch n {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return "Unknown"
	}
}

type Card struct {
	Num  Number
	Suit Suits
}

func (c Card) String() string {
	return c.Num.String() + c.Suit.String()
}

type Deck struct {
	Cards [52]Card
}

func NewDeck() *Deck {
	i := 0
	var n Number = 0
	deck := new(Deck)
	for s := Spade; s <= Club; s++ {
		for n = 1; n <= 13; n++ {
			deck.Cards[i].Num = n
			deck.Cards[i].Suit = s
			i++
		}
	}
	return deck
}
