package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func Reduce[A, B any](items []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, item := range items {
		result = accumulator(result, item)
	}
	return result
}

func Sum[T ~int | ~float64](items []T) T {
	var zero T
	add := func(acc, x T) T { return acc + x }
	return Reduce(items, add, zero)
}

func SumAllTails[T ~int | ~float64](itemsToSum ...[]T) []T {
	var empty []T
	sumTail := func(acc, x []T) []T {
		if len(x) == 0 {
			acc = append(acc, 0)
		} else {
			acc = append(acc, Sum(x[1:]))
		}
		return acc
	}
	return Reduce(itemsToSum, sumTail, empty)
}

func SumAll[T ~int | ~float64](itemsToSum ...[]T) []T {
	var sums []T // SAME AS(?) sums := []T{}
	for _, items := range itemsToSum {
		sums = append(sums, Sum(items))
	}
	return sums
}

func BalanceFor(transactions []Transaction, name string) float64 {
	accrueBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}

		return currentBalance
	}

	return Reduce(transactions, accrueBalance, 0.0)
}

func NewTransaction(from Account, to Account, amount float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: amount}
}

func NewBalanceFor(acc Account, ts []Transaction) Account {
	return Reduce(ts, applyTransaction, acc)
}

func applyTransaction(acc Account, t Transaction) Account {
	if t.From == acc.Name {
		acc.Balance -= t.Sum
	}
	if t.To == acc.Name {
		acc.Balance += t.Sum
	}

	return acc
}
