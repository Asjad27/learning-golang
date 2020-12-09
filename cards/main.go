package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")

	cards1 := newDeckFromFile("my_cards")
	cards1.print()

	cards2 := newDeck()
	cards2.shuffle()
	cards2.print()
}
