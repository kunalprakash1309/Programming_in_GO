package main

import "sort"

type statistics struct {
	numbers []float64
	mean float64
	median float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return stats
}

func sum(numbers []float64) (total float64) {
	for _, v := range numbers {
		total += v
	}
	return
}

func median(numbers []float64) (median float64) {

	if len(numbers) % 2 != 0 {
		median = numbers[len(numbers) + 1]
	} else {
		median = numbers[len(numbers)] + numbers[len(numbers) - 1]
	}

	return median 
}

func main() {

}