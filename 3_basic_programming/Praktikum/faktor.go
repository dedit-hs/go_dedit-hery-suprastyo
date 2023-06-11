package basicProgramming

import "fmt"

func MencariFaktorBilangan(number int) {
	fmt.Printf("Faktor Bilangan %d adalah: ", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}
