package card

import (
	"bank/pkg/bank/types"
)

// IssueCard is method to create new card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

// Withdraw is method to get some amount of money from card
func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active || amount < 0 || amount > card.Balance || amount > 20_000_00 {
		return card
	}

	card.Balance -= amount
	return card
}

// Deposit is method to increase balance in given card by amount
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active || amount < 0 || amount > 50_000_00 {
		return
	}

	card.Balance += amount
}

// AddBonus adds bonus to card balance
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active || card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * percent * daysInMonth / daysInYear / 100

	if bonus > 5_000_00 {
		return
	}

	card.Balance += types.Money(bonus)
}

// Total calculates the total amount of money in all cards.
// Negative balances and inactive cards are ignored.
func Total(cards []types.Card) types.Money {
	var sum types.Money = 0
	for _, card := range cards {
		sum += card.Balance
	}

	return sum
}

// PaymentSources filters cards by balance and activeness
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var result []types.PaymentSource
}
