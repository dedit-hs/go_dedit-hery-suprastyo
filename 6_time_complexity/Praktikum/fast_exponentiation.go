package timeComplexity

func Pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		temp := Pow(x, n/2)
		return temp * temp
	}

	temp := Pow(x, (n-1)/2)
	return x * temp * temp
}
