package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/szalai1/deck-memorizer/pkg/deck"
)

type playOpts struct {
	MaxDeckSize         int      `short:"s" long:"max-deck-size" description:"Maximum number of cards"`
	ShowAssociations    bool     `long:"show-associations" description:"Show associations"`
	CardSelectors       []string `long:"select" description:"select card types. D for diamond, H for heart"`
	AssociationsMapPath string   `long:"associations-file" description:"path to the associations map file"`
}

type Game struct {
	from     deck.Deck
	to       deck.Deck
	recalled deck.Deck
}

type CardAssociation struct {
	Card deck.Card `json:"card"`
	Word string    `json:"word"`
}

func NewSingleSuitGame(s deck.Suit) Game {
	return Game{
		from:     deck.NewSingleSuitDeck(s),
		to:       deck.NewEmptyDeck(),
		recalled: deck.NewEmptyDeck(),
	}
}

func init() {
	parser.AddCommand(
		"memorize",
		"memorize a deck",
		"",
		&playOpts{})
}

func (playOpts *playOpts) Execute(args []string) error {
	return nil
}

func readAssociations(path string) map[deck.Card]string {
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
	return cardWordMapping
}
