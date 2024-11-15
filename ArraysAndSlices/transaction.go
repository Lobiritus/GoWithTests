package ArraysAndSlices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(tran []Transaction, name string) float64 {
	var balance float64
	for _, t := range tran {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
