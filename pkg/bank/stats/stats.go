package stats

import (
	"github.com/Barzu-95/bankapp/pkg/bank/types"
)

//Avg evaluates avarage payment sum
func Avg(payments []types.Payment) types.Money {
	sum := 0

	if len(payments) == 1 {
		return payments[0].Amount
	}
	
	for _, payment := range payments {
		sum = sum + int(payment.Amount)
	}

	avg := types.Money(sum/len(payments)) 

	return avg
}