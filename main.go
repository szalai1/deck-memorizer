package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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

func clearTerminal() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func singleColorGame() {
	deck1 := deck.NewSingleSuitDeck(deck.SuitHeart)
	deck1.Shuffle()
	otherDeck := deck.NewEmptyDeck()
	for c, err := deck1.Draw(); err == nil; c, err = deck1.Draw() {
		fmt.Println(c.String())
		fmt.Scanf("\n")
		clearTerminal()
		otherDeck.PushCardBack(c)
	}
	fmt.Println("Recall")
	for c, err := otherDeck.Draw(); err == nil; c, err = otherDeck.Draw() {
		fmt.Scanf("\n")
		clearTerminal()
		fmt.Println(c.String())
	}
}

func singleSuitGameWithHelp(suit deck.Suit, mapping map[deck.Card]string) {
	deck1 := deck.NewSingleSuitDeck(suit)
	deck1.Shuffle()
	otherDeck := deck.NewEmptyDeck()
	for c, err := deck1.Draw(); err == nil; c, err = deck1.Draw() {
		help, _ := mapping[*c]
		fmt.Println(c.String(), help)
		fmt.Scanf("\n")
		clearTerminal()
		otherDeck.PushCardBack(c)
	}
	fmt.Println("Recall")
	for c, err := otherDeck.Draw(); err == nil; c, err = otherDeck.Draw() {
		fmt.Scanf("\n")
		clearTerminal()
		fmt.Println(c.String())
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
	singleSuitGameWithHelp(deck.SuitDiamond, cardWordMapping)
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
