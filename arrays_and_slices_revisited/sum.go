package main

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
	// sum := zero
	// for _, item := range items {
	// 	sum += item
	// }
	// return sum
}

func SumAll[T ~int | ~float64](itemsToSum ...[]T) []T {
	var sums []T // SAME AS(?) sums := []T{}
	for _, items := range itemsToSum {
		sums = append(sums, Sum(items))
	}
	return sums
}

func SumAllTails[T ~int | ~float64](itemsToSum ...[]T) []T {
	var sums []T
	for _, items := range itemsToSum {
		if len(items) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(items[1:]))
		}
	}
	return sums
}
