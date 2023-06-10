package introductionToAlgorithm

import "fmt"

func CekLampu(number int) {
	if number <= 1 {
		fmt.Println("Lampu mati.")
		return
	}
	lampu := false
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			lampu = !lampu
		}
		fmt.Printf("tombol ke %d lampu %v\n", i, lampu)
	}
	fmt.Printf("kondisi akhir lampu %v\n", lampu)

	if lampu {
		fmt.Println("Lampu menyala.")
	} else {
		fmt.Println("Lampu mati.")
	}
}
