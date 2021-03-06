package array

import "fmt"

func maxProfit(price []int) int {
	// result := 0
	//
	// if len(price) < 2 {
	// 	return result
	// }
	//
	// for i := 0; i < len(price)-1; i++ {
	// 	for j := i + 1; j < len(price); j++ {
	// 		if price[j]-price[i] > result {
	// 			result = price[j] - price[i]
	// 		}
	// 	}
	// }
	// return result

	result := 0

	min := 1 << 32

	fmt.Println("min:", min)

	for i := 0; i < len(price); i++ {
		if price[i] < min {
			min = price[i]
		} else if price[i]-min > result {
			result = price[i] - min
		} else {

		}
	}
	return result

}
