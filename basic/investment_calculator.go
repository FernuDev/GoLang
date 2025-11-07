package main

import (
	"fmt"
	"math"
)

func main() {

	// Data types

	// Integer value
	var investmentAmount = 1000

	// Float value
	var expectedReturnRate = 5.5

	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Printf("Your return value will %f\n", futureValue)

}
