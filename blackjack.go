package blackjack

import(
	"math/rand"
	"time"
	"strconv"
)

//Dec
var SUITS = []string{"♠", "♥", "♣", "♦"}
var FACES = []string{"J", "Q", "K", "A"}

type Hand struct {
	Cards []*Card
	Total int
}
type Card struct {
	Symbol string
	Suit string
	Weight int
}

//Init
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Util
func getCard() *Card {
	newCard := new(Card)
	newCard.Suit = SUITS[rand.Intn(3)]
	newCard.Weight = rand.Intn(10-1) + 1
	
	if newCard.Weight > 10 {
		newCard.Symbol = FACES[rand.Intn(2)]
		newCard.Weight = 10 //Switch weight back to 10
	} else if newCard.Weight > 1 && newCard.Weight < 10 {
		newCard.Symbol = strconv.Itoa(newCard.Weight)
	} else {
		newCard.Symbol = FACES[3]
		newCard.Weight = 11 //Starts as 11 and is set lower if total gets over 21
	}
	
	return newCard
}

func GetHand() *Hand {
	newHand := new(Hand)
	newCards := []*Card{getCard(), getCard()}
	newHand.Cards = newCards
	newHand.Total = getHandTotal(newHand)
	
	//Has 2 Aces
	if newHand.Total == 22 {
		newHand.Cards[0].Weight = 1
	}
	
	return newHand
}

func AddCard(currentHand *Hand) *Hand {
	newHand := currentHand
	newHand.Cards = append(currentHand.Cards, getCard())
	newHand.Total = getHandTotal(newHand)
	return newHand
}

func CheckHand(currentHand *Hand) int {
	newHand := currentHand
	checked := false
	
	//0 = under, 1 = Win Condition, 2 = Bust	
	if newHand.Total > 21 {
		//Ace complicatedness
		checked, newHand = checkAces(currentHand)
		if !checked {
			return 2
		}	
	} else if newHand.Total == 21 {
		return 1
	} else {
		return 0
	}
	
	return CheckHand(newHand)
}

func checkAces(currentHand *Hand) (bool, *Hand) {
	newHand := currentHand
	for _, element := range newHand.Cards {
		if element.Symbol == "A" && element.Weight != 1 {
			element.Weight = 1
			newHand.Total = getHandTotal(newHand)
			if getHandTotal(newHand) < 21 {
				return true, newHand
			}
		}
	}
	return false, newHand
}

func getHandTotal(currentHand *Hand) int {
	total := 0
	for _, element := range currentHand.Cards {
		total += element.Weight
	}
	return total
}