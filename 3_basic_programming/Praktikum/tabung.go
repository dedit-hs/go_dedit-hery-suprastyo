package basicProgramming

func HitungLuasPermukaanTabung(r, T float64) float64 {
	pi := 3.14

	L := 2 * pi * r * (r + T)

	return L
}
