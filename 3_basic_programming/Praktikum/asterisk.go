package basicProgramming

import "fmt"

func PlayWithAsterisk(number int) {
	for i := 0; i < number; i++ {
		for j := 0; j < number-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
