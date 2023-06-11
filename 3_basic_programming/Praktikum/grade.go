package basicProgramming

import "fmt"

func HitungGradeNilai(nilai int) {
	if nilai >= 80 && nilai <= 100 {
		fmt.Printf("Nilai %d mendapat grade A", nilai)
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Printf("Nilai %d mendapat grade B", nilai)
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Printf("Nilai %d mendapat grade C", nilai)
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Printf("Nilai %d mendapat grade D", nilai)
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Printf("Nilai %d mendapat grade E", nilai)
	} else {
		fmt.Printf("Nilai %d adalah Nilai Invalid", nilai)
	}
	fmt.Println("")
}
