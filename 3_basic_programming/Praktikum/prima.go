package basicProgramming

import "fmt"

func CekBilanganPrima(number int) {
	if number >= 2 {
		for i := 2; i < number-1; i++ {
			if number%i == 0 {
				fmt.Printf("%d bukan bilangan prima\n", number)
				return
			}
		}
		fmt.Printf("%d merupakan bilangan prima\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", number)
	}
}
