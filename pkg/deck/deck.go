package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	SuitHeart = iota
	SuitDiamond
	SuitClub
	SuitSpade
)

const (
	_ = iota
	RankAce
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	Rank9
	Rank10
	RankJack
	RankQueen
	RankKing
)

type Suit int
type Rank int

type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
}

type Deck interface {
	Shuffle()
	Size() int
	Draw() (*Card, error)
	PushCardBack(c ...*Card)
	PushCardFront(c ...*Card)
}

func NewFullOrderedDeck() Deck {
	d := deck{cards: make([]*Card, 0, 52)}
	for _, s := range []Suit{SuitHeart, SuitDiamond, SuitSpade, SuitClub} {
		for _, r := range []Rank{
			RankAce, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8, Rank9, Rank10, RankJack, RankQueen, RankKing} {
			card := &Card{Suit: s, Rank: r}
			d.cards = append(d.cards, card)
		}
	}
	return &d
}

func NewEmptyDeck() Deck {
	return &deck{cards: make([]*Card, 0)}
}

func NewSingleSuitDeck(s Suit) Deck {
	d := &deck{cards: make([]*Card, 0, 13)}
	for _, r := range []Rank{
		RankAce, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8, Rank9, Rank10, RankJack, RankQueen, RankKing} {
		card := &Card{Suit: s, Rank: r}
		d.cards = append(d.cards, card)
	}
	return d
}

func NewCardFromString(input string) (*Card, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return nil, errors.New("card format is invalid")
	}
	parts[0] = strings.ToUpper(parts[0])
	var card Card
	switch parts[0] {
	case "H":
		card.Suit = SuitHeart
	case "D":
		card.Suit = SuitDiamond
	case "S":
		card.Suit = SuitSpade
	case "C":
		card.Suit = SuitClub
	default:
		return nil, fmt.Errorf("could not recognize %s as suit", parts[0])
	}
	switch parts[1] {
	case "A":
		card.Rank = RankAce
		return &card, nil
	case "J":
		card.Rank = RankJack
		return &card, nil
	case "Q":
		card.Rank = RankQueen
		return &card, nil
	case "K":
		card.Rank = RankKing
		return &card, nil
	default:
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("could not recognize %s as rank", parts[1])
		}
		if num < 2 || num > 10 {
			return nil, fmt.Errorf("%s is not a valid rank", parts[1])
		}
		card.Rank = Rank(num)
		return &card, nil
	}
}

func (c *Card) String() string {
	suitM := map[Suit]string{
		SuitHeart:   "heart",
		SuitDiamond: "diamond",
		SuitSpade:   "spade",
		SuitClub:    "club",
	}
	rankM := map[Rank]string{
		RankAce:   "a",
		Rank2:     "2",
		Rank3:     "3",
		Rank4:     "4",
		Rank5:     "5",
		Rank6:     "6",
		Rank7:     "7",
		Rank8:     "8",
		Rank9:     "9",
		Rank10:    "10",
		RankJack:  "j",
		RankQueen: "q",
		RankKing:  "k",
	}
	return fmt.Sprintf("%s-%s", suitM[c.Suit], rankM[c.Rank])
}

func (c *Card) absoluteRank() int {
	return int(c.Suit)*13 + int(c.Rank)
}

type deck struct {
	cards []*Card
}

func (deck *deck) Shuffle() {
	deck.fisherYatesShuffle()
}

func (deck *deck) fisherYatesShuffle() {
	if deck.Size() < 2 {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := deck.Size() - 1; i > 0; i-- {
		n := rand.Intn(i + 1)
		tmp := deck.cards[n]
		deck.cards[n] = deck.cards[i]
		deck.cards[i] = tmp
	}
}

func (deck *deck) Size() int {
	return len(deck.cards)
}

func (deck *deck) Draw() (*Card, error) {
	if deck.Size() == 0 {
		return nil, errors.New("empty deck")
	}
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card, nil
}

func (deck *deck) PushCardBack(c ...*Card) {
	deck.cards = append(deck.cards, c...)
}

func (deck *deck) PushCardFront(cards ...*Card) {
	newCards := make([]*Card, 0, len(cards)+deck.Size())
	for _, c := range cards {
		newCards = append(newCards, c)
	}
	newCards = append(newCards, deck.cards...)
	deck.cards = newCards
}
