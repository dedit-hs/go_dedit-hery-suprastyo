package main

import (
	"fmt"

	introductionToAlgorithm "github.com/dedit-hs/go_dedit-hery-suprastyo/1_introduction_to_algorithm/Praktikum"
	versionControlAndBranchManagement "github.com/dedit-hs/go_dedit-hery-suprastyo/2_version_control_and_branch_management_git/Praktikum"
	basicProgramming "github.com/dedit-hs/go_dedit-hery-suprastyo/3_basic_programming/Praktikum"
	dataStructure "github.com/dedit-hs/go_dedit-hery-suprastyo/4_data_structure/Praktikum"
)

func main() {
	fmt.Println("***** Introduction to Algorithm *****")
	sectionIntroToAlgorithm() // Tugas Praktikum 1
	fmt.Println("***** Version Control and Branch Management *****")
	sectionVersionControl() // Tugas Praktikum 2
	fmt.Println("***** Basic Programming *****")
	sectionBasicProgramming() // Tugas Praktikum 3
	fmt.Println("***** Data Structure *****")
	sectionDataStructure() // Tugas Praktikum 4
}

func sectionIntroToAlgorithm() {
	n := 5
	fmt.Println("=== Cek Bilangan Prima ===")
	introductionToAlgorithm.IsPrima(n)

	fmt.Println("=== Cek Kondisi Lampu ===")
	introductionToAlgorithm.CekLampu(n)

	fmt.Println("=== Hello ===")
	introductionToAlgorithm.Hello()
}

func sectionVersionControl() {
	versionControlAndBranchManagement.UpdateFeatureA()
	versionControlAndBranchManagement.FeatureB()
	fmt.Println("add new featureA")
	fmt.Println("update featureA")
	fmt.Println("new update featureA")
}

func sectionBasicProgramming() {
	fmt.Println("=== Hitung Luas Permukaan Tabung ===")
	tabung := basicProgramming.HitungLuasPermukaanTabung(4, 20)
	fmt.Println("Luas tabung: ", tabung)
	fmt.Println("=== Hitung Grade Nilai ===")
	basicProgramming.HitungGradeNilai(10)
	fmt.Println("=== Mencari Faktor Bilangan ===")
	basicProgramming.MencariFaktorBilangan(20)
	fmt.Println("=== Cek Bilangan Prima ===")
	basicProgramming.CekBilanganPrima(11)
	fmt.Println("=== Cek Palindrome ===")
	fmt.Println(basicProgramming.Palindrome("rusak"))
	fmt.Println("=== Exponentiation ===")
	fmt.Println(basicProgramming.HitungPangkat(5, 5))
	fmt.Println("=== Play with Asterisk ===")
	basicProgramming.PlayWithAsterisk(5)
	fmt.Println("=== Tabel Perkalian ===")
	basicProgramming.TabelPerkalian(9)
}

func sectionDataStructure() {
	fmt.Println(dataStructure.Merge([]string{"saya", "dia", "mereka"}, []string{"kita", "saya"}))
	fmt.Println(dataStructure.MunculSekali("223345566789"))
	fmt.Println(dataStructure.PairSum([]int{2, 5, 9, 11}, 11))
}
