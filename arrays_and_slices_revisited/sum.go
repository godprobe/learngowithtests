package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func Reduce[A any](items []A, accumulator func(A, A) A, initialValue A) A {
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

func BalanceFor(transations []Transaction, name string) float64 {
	return 0.0
}
