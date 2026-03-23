package main

func FinalPrices(prices []int) []int {

	answers := []int{}
	for i := range prices {

		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				discount = prices[j]
				break
			}
		}

		answers = append(answers, prices[i]-discount)

	}
	return answers
}
