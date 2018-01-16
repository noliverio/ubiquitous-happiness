package main

import "fmt"

const (
	//Create a huge number by bit shifting 1 to the left 100 places

	Big = 1 << 100

	// change the big number into 2 via bit-shifting again

	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
