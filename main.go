package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/szalai1/deck-memorizer/pkg/deck"
)

type CardAssociation struct {
	Card deck.Card `json:"card"`
	Word string    `json:"word"`
}

type Game struct {
	from  deck.Deck
	to    deck.Deck
	score int
}

func NewGame() Game {
	return Game{
		from:  deck.NewSingleSuitDeck(deck.SuitDiamond),
		to:    deck.NewEmptyDeck(),
		score: 0,
	}
}

func main() {
	var cardWordPairs []CardAssociation
	file, err := ioutil.ReadFile("mapping.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &cardWordPairs)
	if err != nil {
		log.Fatal(err)
	}
	cardWordMapping := map[deck.Card]string{}
	for _, p := range cardWordPairs {
		cardWordMapping[p.Card] = p.Word
	}
	fmt.Println(cardWordMapping)
	/*heart := deck.NewSingleSuitDeck(deck.SuitHeart)
	otherDeck := deck.NewEmptyDeck()
	heart.Shuffle()
	for c, err := heart.Draw(); err == nil; c, err = heart.Draw() {
		fmt.Println(c)
		otherDeck.PushCardBack(c)
		//fmt.Scanf("\n")
	}
	*/
}
