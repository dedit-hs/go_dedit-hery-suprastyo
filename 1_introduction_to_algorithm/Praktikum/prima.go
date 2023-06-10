package introductionToAlgorithm

import "fmt"

func IsPrima(number int) {
	if number > 1 && number <= 50 {
		for i := 2; i < number-1; i++ {
			if number%i == 0 {
				fmt.Printf("%d adalah bukan bilangan prima\n", number)
				return
			}
		}
		fmt.Printf("%d adalah bilangan prima\n", number)
	} else {
		fmt.Printf("%d adalah bukan bilangan prima\n", number)
	}
}
