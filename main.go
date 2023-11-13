package main

func main() {

	cards := deck{"Ace of diamond", newCard()}

	cards = append(cards, "six of Spades")

	cards.print()
}

func newCard() string {
	return "i m a new card"
}