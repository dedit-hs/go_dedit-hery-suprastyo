package dataStructure

import "strconv"

func MunculSekali(angka string) []int {
	var uniqueAngka = make(map[int]int)
	for _, value := range angka {
		if digit, err := strconv.Atoi(string(value)); err == nil {
			uniqueAngka[digit]++
		}
	}
	sekali := make([]int, 0)
	for key, value := range uniqueAngka {
		if value == 1 {
			sekali = append(sekali, key)
		}
	}
	return sekali
}
