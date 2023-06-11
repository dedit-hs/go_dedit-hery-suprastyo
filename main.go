package main

import (
	"fmt"

	introductionToAlgorithm "github.com/dedit-hs/go_dedit-hery-suprastyo/1_introduction_to_algorithm/Praktikum"
	versionControlAndBranchManagement "github.com/dedit-hs/go_dedit-hery-suprastyo/2_version_control_and_branch_management_git/Praktikum"
)

func main() {
	sectionIntroToAlgorithm() // Tugas Praktikum 1
	sectionVersionControl()   // Tugas Praktikum 2
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

func sectionVersionControl() {
	versionControlAndBranchManagement.UpdateFeatureA()
	fmt.Println("add new featureA")
}
