package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 5_000_00, Active: true}, 7_000_00)
	fmt.Println(result.Balance)
	// Output: 500000
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleDeposit_normal() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_inactiveCard() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_normal() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 2, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2001643
}

func ExampleAddBonus_inactiveCard() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 2, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -20_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 2, 30, 365)
	fmt.Println(card.Balance)
	// Output: -2000000
}

func ExampleAddBonus_tooBigBonus() {
	card := types.Card{Balance: 200_000_00, MinBalance: 100_000_00, Active: true}
	AddBonus(&card, 99, 30, 365)
	fmt.Println(card.Balance)
	// Output: 20000000
}

func ExampleTotal_normal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	//Output: 2000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: 999_99,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 888_88,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: -1,
			Active:  true,
		},
	}

	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

	//Output: 5058 xxxx xxxx 8888

}
