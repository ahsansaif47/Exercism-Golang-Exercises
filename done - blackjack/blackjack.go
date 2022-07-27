package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, d_card := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	tot := c1 + c2
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case tot == 21 && d_card < 10:
		return "W"
	case tot == 21 && (d_card == 10 || d_card == 11):
		return "S"
	case tot >= 17 && tot <= 20:
		return "S"
	case tot >= 12 && tot <= 16:
		if d_card < 7 {
			return "S"
		} else {
			return "H"
		}
	case (c1 + c2) <= 11:
		return "H"
	}
	return ""
}
