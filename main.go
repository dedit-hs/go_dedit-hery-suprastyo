package main

import (
	"fmt"

	introductionToAlgorithm "github.com/dedit-hs/go_dedit-hery-suprastyo/1_introduction_to_algorithm/Praktikum"
)

func main() {
	sectionIntroToAlgorithm() // Tugas Praktikum 1
}

func sectionIntroToAlgorithm() {
	n := 5
	fmt.Println("===== Cek Bilangan Prima =====")
	introductionToAlgorithm.IsPrima(n)

	fmt.Println("===== Cek Kondisi Lampu =====")
	introductionToAlgorithm.CekLampu(n)

	fmt.Println("===== Hello =====")
	introductionToAlgorithm.Hello()
}
