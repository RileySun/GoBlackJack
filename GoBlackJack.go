package main

import(
	"math/rand"
	"time"
	"strconv"
)

//Dec
var SUITS = []string{"♠", "♥", "♣", "♦"}
var FACES = []string{"J", "Q", "K"}

type Hand struct {
	Cards []*Card
	Total int
}
type Card struct {
	Symbol string
	Suit string
	Weight int
}

//Main

func main() {
	rand.Seed(time.Now().UnixNano())
}

//Util

func getCard() *Card {
	//Weights are counted to 13 to get face cards, then switched back to 10 to follow blackjack rules
	newCard := new(Card)
	newCard.Suit = SUITS[rand.Intn(3)]
	newCard.Weight = rand.Intn(13-1) + 1
	
	if newCard.Weight > 10 {
		newCard.Symbol = FACES[rand.Intn(2)]
		newCard.Weight = 10 //Switch weight back to 10
	} else {
		newCard.Symbol = strconv.Itoa(newCard.Weight)
	}
	
	return newCard
}

func GetHand() *Hand {
	newHand := new(Hand)
	newCards := []*Card{getCard(), getCard()}
	newHand.Cards = newCards
	newHand.Total = newHand.Cards[0].Weight + newHand.Cards[1].Weight
	
	return newHand
}

func AddCard(currentHand *Hand) *Hand {
	newCards := append(currentHand.Cards, getCard())
	currentHand.Cards = newCards
	
	total := 0
	for _, element := range currentHand.Cards {
		total += element.Weight
	}
	currentHand.Total = total
	
	return currentHand
}

func CheckHand(currentHand *Hand) int {
	//0 = under, 1 = Win Condition, 2 = Bust
	if currentHand.Total < 21 {
		return 0
	} else if currentHand.Total == 21 {
		return 1
	} else {
		return 2
	}
}