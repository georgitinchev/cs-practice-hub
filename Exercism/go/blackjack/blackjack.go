package blackjack


func ParseCard(card string) int {
	switch card {
	case "A", "ace":
		return 11
	case "K", "Q", "J", "king", "queen", "jack":
		return 10
	case "10", "ten":
		return 10
	case "9", "nine":
		return 9
	case "8", "eight":
		return 8
	case "7", "seven":
		return 7
	case "6", "six":
		return 6
	case "5", "five":
		return 5
	case "4", "four":
		return 4
	case "3", "three":
		return 3
	case "2", "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	   playerTotal := ParseCard(card1) + ParseCard(card2)
	   dealerValue := ParseCard(dealerCard)

	   // Pair of aces
	   if (card1 == "A" || card1 == "ace") && (card2 == "A" || card2 == "ace") {
			   return "P"
	   }
	   // Pair of 10-point cards (jack, queen, king, ten)
	   tenCards := map[string]bool{"10": true, "ten": true, "J": true, "jack": true, "Q": true, "queen": true, "K": true, "king": true}
	   if tenCards[card1] && tenCards[card2] {
			   return "S"
	   }
	   // Pair of twos or fives
	   if (card1 == "2" || card1 == "two") && (card2 == "2" || card2 == "two") {
			   return "H"
	   }
	   if (card1 == "5" || card1 == "five") && (card2 == "5" || card2 == "five") {
			   return "H"
	   }
	   // Blackjack (ace + 10-point card)
	   if (ParseCard(card1) == 11 && ParseCard(card2) == 10) || (ParseCard(card1) == 10 && ParseCard(card2) == 11) {
			   if dealerValue == 11 || dealerValue == 10 {
					   return "S"
			   } else {
					   return "W"
			   }
	   }
	   // All other cases
	   if playerTotal >= 17 {
			   return "S"
	   } else if playerTotal >= 12 && dealerValue < 7 {
			   return "S"
	   } else {
			   return "H"
	   }
}
