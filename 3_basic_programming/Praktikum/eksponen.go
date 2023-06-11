package basicProgramming

func HitungPangkat(angka, pangkat int) int {
	var hasil int = angka
	for i := 1; i < pangkat; i++ {
		hasil *= angka
	}
	return hasil
}
