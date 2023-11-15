package main

func main() {

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// hand, remainingcards := deal(cards, 5)
	// hand.print()
	// remainingcards.print()

	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.schuffle()
	cards.print()
}