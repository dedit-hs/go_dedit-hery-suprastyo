package main

import (
	"fmt"

	introductionToAlgorithm "github.com/dedit-hs/go_dedit-hery-suprastyo/1_introduction_to_algorithm/Praktikum"
	versionControlAndBranchManagement "github.com/dedit-hs/go_dedit-hery-suprastyo/2_version_control_and_branch_management_git/Praktikum"
	basicProgramming "github.com/dedit-hs/go_dedit-hery-suprastyo/3_basic_programming/Praktikum"
	dataStructure "github.com/dedit-hs/go_dedit-hery-suprastyo/4_data_structure/Praktikum"
	pointerAndErrorHandling "github.com/dedit-hs/go_dedit-hery-suprastyo/5_pointer_and_error_handling/Praktikum"
	timeComplexity "github.com/dedit-hs/go_dedit-hery-suprastyo/6_time_complexity/Praktikum"
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
	fmt.Println("***** Pointer, Advanced Function, Struct Etc. *****")
	sectionPointerAndErrorHandling() // Tugas Praktikum 5
	fmt.Println("\n***** Time Complexity *****")
	sectionTimeComplexity() // Tugas Praktikum 6
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

func sectionPointerAndErrorHandling() {
	fmt.Println(pointerAndErrorHandling.CompareString("LEMFOXBRO", "FOX"))
	fmt.Println(pointerAndErrorHandling.Caesar(2, "alta"))
	a := 12
	b := 3
	pointerAndErrorHandling.Swap(&a, &b)
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Printf("Masukkan nilai 1: ")
	fmt.Scan(&a1)
	fmt.Printf("Masukkan nilai 2: ")
	fmt.Scan(&a2)
	fmt.Printf("Masukkan nilai 3: ")
	fmt.Scan(&a3)
	fmt.Printf("Masukkan nilai 4: ")
	fmt.Scan(&a4)
	fmt.Printf("Masukkan nilai 5: ")
	fmt.Scan(&a5)
	fmt.Printf("Masukkan nilai 6: ")
	fmt.Scan(&a6)
	min, max = pointerAndErrorHandling.GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min: ", min)
	fmt.Println("Nilai Max: ", max)

	var std = pointerAndErrorHandling.Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Printf("Input Student's %d Name : ", i)
		fmt.Scan(&name)
		std.Name = append(std.Name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		std.Score = append(std.Score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", std.Average())
	scoreMax, nameMax := std.Max()
	fmt.Println("Max score Students is: "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := std.Min()
	fmt.Println("Min score Students is: "+nameMin+" (", scoreMin, ")")

	var menu int
	var siswa = pointerAndErrorHandling.Siswa{}
	var c pointerAndErrorHandling.Cipher = &siswa
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu?")
	fmt.Scan(&menu)
	switch menu {
	case 1:
		fmt.Print("\nInput Nama Siswa: ")
		fmt.Scan(&siswa.Nama)
		fmt.Print("\nEncode dari Nama Siswa " + siswa.Nama + " is : " + c.Encode())
	case 2:
		fmt.Print("\nInput Nama Siswa: ")
		fmt.Scan(&siswa.Nama)
		fmt.Print("\nDecode dari Nama Siswa " + siswa.Nama + " is : " + c.Decode())
	default:
		fmt.Println("Wrong input name menu.")
	}
}

func sectionTimeComplexity() {
	n := 12
	isPrime := timeComplexity.PrimeNumber(n)

	if isPrime {
		fmt.Printf("%d adalah bilangan prima.\n", n)
	} else {
		fmt.Printf("%d BUKAN bilangan prima.\n", n)
	}

	fmt.Println(timeComplexity.Pow(2, 10))
}
