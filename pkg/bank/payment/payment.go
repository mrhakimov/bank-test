package payment

import "bank/pkg/bank/types"

// Max returns payment with max amount
func Max(payments []types.Payment) types.Payment {
	var maxAmount int = int(payments[0].Amount)
	var maxIndex int = 0

	for i, payment := range payments {
		if maxAmount < int(payment.Amount) {
			maxAmount = int(payment.Amount)
			maxIndex = i
		}
	}

	return payments[maxIndex]
}
