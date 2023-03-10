/*
In this exercise we will simulate the first turn of a Blackjack game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:

card	value	card	value
ace	11	eight	8
two	2	nine	9
three	3	ten	10
four	4	jack	10
five	5	queen	10
six	6	king	10
seven	7	other	0
Note: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)
Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

Category: Large Hand

If you have a pair of aces you must always split them.
If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
Category: Small Hand

If your cards sum up to 17 or higher you should always stand.
If your cards sum up to 11 or lower you should always hit.
If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
The overall logic has already been implemented. You have four tasks:

1. Calculate the score of any given card.
Implement a function to calculate the numerical value of a card given its name using conditionals.

value := ParseCard("ace")
fmt.Println(value)
// Output: 11
2. Determine if two cards make up a Blackjack.
Implement a function that returns true if two cards form a Blackjack, false otherwise.

isBlackjack := IsBlackjack("queen", "ace")
fmt.Println(isBlackjack)
// Output: true
3. Implement the decision logic for hand scores larger than 20 points.
Implement a function that returns the string representation of a decision given your cards. This function is only called if the handScore is larger than 20. It will receive 2 arguments: isBlackJack and dealerScore. It should implement the bulletpoints in the category "Large Hand" above.

isBlackJack := true
dealerScore := 7
choice := LargeHand(isBlackJack, dealerScore)
fmt.Println(choice)
// Output: "W"
4. Implement the decision logic for hand scores with less than 21 points.
Implement a function that returns the string representation of a decision given your cards. This function is only called if the handScore is less than 21. It will receive 2 arguments: handScore and dealerScore. It should implement the bulletpoints in the category "Small Hand" above.

handScore := 15
dealerScore := 12
choice := SmallHand(handScore, dealerScore)
fmt.Println(choice)
// Output: "H"
*/

package main

func GetCardValue(cardName string) int {
	switch cardName {
	case "ace":
		return 11
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func IsBlackjack(card1, card2 string) bool {
	return GetCardValue(card1)+GetCardValue(card2) == 21
}

func LargeHand(isBlackJack bool, dealerScore int) string {

	// 21
	if isBlackJack {
		if dealerScore < 10 || dealerScore > 11 {
			return "W"
		} else {
			return "S"
		}
	} else {
		return "P"
	}
}

func SmallHand(handScore, dealerScore int) string {

	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else if handScore >= 12 && handScore <= 16 {
		if dealerScore >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
	return ""
}

// concString:=strings.Join(slice...,"<what_you_want_between_>")
//{"ast","asf","atgawet"}
// concString:=strings.Join(slice...,"-")  // ast-asf-atgawet

// stplittedString:=strings.Split(<string>, <threshold>)
// myString:= "my name is tarun"
// stplittedString:=strings.Split(myString, " ") // {my, name, is, tarun}