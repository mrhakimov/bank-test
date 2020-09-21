package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax_normal() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 10,
		},
		{
			ID:     99,
			Amount: 200,
		},
		{
			ID:     100,
			Amount: 200,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 99
}
