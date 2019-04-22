package deck

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewFullOrderedDeck()
	if d.Size() != 52 {
		t.Error("default deck should have 52 cards")
	}
	for _, s := range []Suit{SuitHeart, SuitDimond, SuitSpade, SuitClub} {
		for _, r := range []Rank{
			RankAce, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8, Rank9, Rank10, RankJack, RankQueen, RankKing} {
			expectedCard := Card{Suit: s, Rank: r}
			nextCard, err := d.Draw()
			if err != nil {
				t.Errorf("the deck should not be empty")
			}
			if *nextCard != expectedCard {
				t.Errorf("%s != %s", &expectedCard, nextCard)
			}
		}
	}
	if d.Size() != 0 {
		t.Error("deck should be empty")
	}
}

func TestCardString(t *testing.T) {
	d := NewFullOrderedDeck()
	c, err := d.Draw()
	if err != nil {
		t.Fail()
	}
	if c.String() != "Heart-Ace" {
		t.Errorf("Got: %s, should be: %s", c.String(), "Heart - Ace")
	}
}

func TestShuffle(t *testing.T) {
	// TODO measure randomness
	d := NewFullOrderedDeck()
	d.Shuffle()
	for c, err := d.Draw(); err == nil; c, err = d.Draw() {
		fmt.Printf("%s ", c)
	}
}

func TestPushCardBack(t *testing.T) {
	d := NewEmptyDeck()
	d.PushCardBack(&Card{Suit: SuitClub, Rank: Rank10})
	if d.Size() != 1 {
		t.Errorf("size should be 1 instead of %d", d.Size())
	}
	d.PushCardBack(&Card{Suit: SuitDimond, Rank: Rank10})
	d.PushCardBack(&Card{Suit: SuitHeart, Rank: Rank10})
	c, err := d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitClub || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}
	c, err = d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitDimond || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}

	c, err = d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitHeart || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}
	_, err = d.Draw()
	if err == nil {
		t.Error("deck should be empty")
	}
}

func TestPushCardFront(t *testing.T) {
	d := NewEmptyDeck()
	d.PushCardFront(&Card{Suit: SuitClub, Rank: Rank10})
	if d.Size() != 1 {
		t.Errorf("size should be 1 instead of %d", d.Size())
	}
	d.PushCardFront(&Card{Suit: SuitDimond, Rank: Rank10})
	d.PushCardFront(&Card{Suit: SuitHeart, Rank: Rank10})
	c, err := d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitHeart || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}
	c, err = d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitDimond || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}

	c, err = d.Draw()
	if err != nil {
		t.Error(err)
	}
	if c.Suit != SuitClub || c.Rank != Rank10 {
		t.Errorf(" card is not Cluber-10 but %s", c)
	}
	_, err = d.Draw()
	if err == nil {
		t.Error("deck should be empty")
	}
}
