package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/szalai1/deck-memorizer/pkg/deck"
)

type selector func(deck.Card) bool

type playOpts struct {
	MaxDeckSize         int      `short:"s" long:"max-deck-size" description:"Maximum number of cards"`
	ShowAssociations    bool     `long:"show-associations" description:"Show associations"`
	CardSelectors       []string `long:"select" description:"select card types. D for diamond, H for heart"`
	AssociationsMapPath string   `long:"associations-file" description:"path to the associations map file"`
}

type Game struct {
	from            deck.Deck
	to              deck.Deck
	recalled        deck.Deck
	learnTime       time.Duration
	config          *playOpts
	associationsMap map[deck.Card]string
}

type CardAssociation struct {
	Card deck.Card `json:"card"`
	Word string    `json:"word"`
}

func (g *Game) Play() {
	// show
	for c, err := g.from.Draw(); err != nil; c, err = g.from.Draw() {
		time.Sleep(g.learnTime)
		if ass, ok := g.associationsMap[*c]; g.config.ShowAssociations && ok {
			fmt.Printf("%v %s", c, ass)
		} else {
			fmt.Printf("%v", c)
		}
		g.to.PushCardBack(c)
	}
	// recall
	for c, err := g.to.Draw(); err != nil; c, err = g.to.Draw() {
		card := readCardFromStdIn()
		g.recalled.PushCardFront(card)
		g.from.PushCardFront(c)
	}
}

func readCardFromStdIn() *deck.Card {
	for {
		var cardStr string
		fmt.Scanf("%s", &cardStr)
		c, err := deck.NewCardFromString(cardStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return c
	}
}

func init() {
	parser.AddCommand(
		"memorize",
		"memorize a deck",
		"",
		&playOpts{})
}

func (config *playOpts) NewGame(s selector) Game {
	fullDeck := deck.NewFullOrderedDeck()
	g := Game{
		from:            deck.NewEmptyDeck(),
		to:              deck.NewEmptyDeck(),
		recalled:        deck.NewEmptyDeck(),
		associationsMap: readAssociations(config.AssociationsMapPath),
	}
	for c, err := fullDeck.Draw(); err != nil; c, err = fullDeck.Draw() {
		if s(*c) {
			g.from.PushCardFront(c)
		}
	}
	g.from.Shuffle()
	return g
}

func (playOpts *playOpts) Execute(args []string) error {
	selector := falseSelector
	for _, s := range playOpts.CardSelectors {
		newSelector, err := newSelectorFromString(s)
		if err != nil {
			return fmt.Errorf("memorize execute failed: %s", err)
		}
		selector = orSelector(selector, newSelector)
	}
	g := playOpts.NewGame(selector)
	g.Play()
	return nil
}

func falseSelector(c deck.Card) bool {
	return false
}

func newSelectorFromString(selector string) (selector, error) {
	var suit deck.Suit
	switch selector {
	case "D":
		suit = deck.SuitDiamond
	case "H":
		suit = deck.SuitHeart
	case "S":
		suit = deck.SuitSpade
	case "C":
		suit = deck.SuitClub
	default:
		return nil, fmt.Errorf("could not parse selector: %s ")
	}
	return func(c deck.Card) bool {
		return c.Suit == suit
	}, nil
}

func orSelector(selectorA selector, selectorB selector) selector {
	return func(c deck.Card) bool {
		return selectorA(c) || selectorB(c)
	}
}

func readAssociations(path string) map[deck.Card]string {
	var cardWordPairs []CardAssociation
	file, err := ioutil.ReadFile("mapping.json")
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal([]byte(file), &cardWordPairs)
	if err != nil {
		log.Print(err)
		return cardWordPairs
	}
	cardWordMapping := map[deck.Card]string{}
	for _, p := range cardWordPairs {
		cardWordMapping[p.Card] = p.Word
	}
	return cardWordMapping
}
