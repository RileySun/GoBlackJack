package blackjack

import "testing"

func TestAces(t *testing.T) {
    hand := GetHand()
	
	hand.Cards[0].Symbol = "A"
	hand.Cards[0].Weight = 11
	
	hand.Cards[1].Symbol = "A"
	hand.Cards[1].Weight = 11
	
	hand.Total = getHandTotal(hand)
	
	hand = AddCard(hand)
	
    result := CheckHand(hand)
    
    if result != 0 && result != 1 {
       t.Errorf("Total was over 21, got: %d, want: under 21.", hand.Total)
    }
}

func TestOver(t *testing.T) {
    hand := GetHand()
	hand = AddCard(hand)
	
	hand.Cards[0].Symbol = "Q"
	hand.Cards[0].Weight = 10
	
	hand.Cards[1].Symbol = "Q"
	hand.Cards[1].Weight = 10
	
	hand.Cards[2].Symbol = "Q"
	hand.Cards[2].Weight = 10
	
	hand.Total = getHandTotal(hand)
	
    result := CheckHand(hand)
    
    if result != 2 {
       t.Errorf("Total was under 21, got: %d, want: over 21.", hand.Total)
    }
}

func TestUnder(t *testing.T) {
    hand := GetHand()
	hand = AddCard(hand)
	
	hand.Cards[0].Symbol = "2"
	hand.Cards[0].Weight = 2
	
	hand.Cards[1].Symbol = "2"
	hand.Cards[1].Weight = 2
	
	hand.Cards[2].Symbol = "2"
	hand.Cards[2].Weight = 2
	
	hand.Total = getHandTotal(hand)
	
    result := CheckHand(hand)
    
    if result != 0 {
       t.Errorf("Total was over 21, got: %d, want: under 21.", hand.Total)
    }
}

func TestWin(t *testing.T) {
    hand := GetHand()
	hand = AddCard(hand)
	
	hand.Cards[0].Symbol = "A"
	hand.Cards[0].Weight = 11
	
	hand.Cards[1].Symbol = "5"
	hand.Cards[1].Weight = 5
	
	hand.Cards[2].Symbol = "5"
	hand.Cards[2].Weight = 5
	
	hand.Total = getHandTotal(hand)
	
	
    result := CheckHand(hand)
    
    if result != 1 {
       t.Errorf("Total was not 21, got: %d, want: 21.", hand.Total)
    }
}

func TestBlackjack(t *testing.T) {
    hand := GetHand()
	
	hand.Cards[0].Symbol = "A"
	hand.Cards[0].Weight = 11
	
	hand.Cards[1].Symbol = "Q"
	hand.Cards[1].Weight = 10
	
	hand.Total = getHandTotal(hand)
	
    result := CheckHand(hand)
    
    if result != 1 {
       t.Errorf("Total was not 21, got: %d, want: 21.", hand.Total)
    }
}