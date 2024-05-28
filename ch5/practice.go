package main

import "fmt"

func main() {
	title_5_1()
	title_5_2()
	title_5_3()
}

func title_5_1() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func title_5_2() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}

func title_5_3() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
