package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var investmentAmount, years, expectedReturnRate float64

	// Input
	fmt.Print("Enter the investment value: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the years for the investment: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expectated return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	furuteRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Your return value will %f\n", futureValue)
	fmt.Printf("Your return value in real terms will %f\n", furuteRealValue)

}
