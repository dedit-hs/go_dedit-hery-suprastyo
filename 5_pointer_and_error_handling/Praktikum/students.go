package pointerAndErrorHandling

type Student struct {
	Name  []string
	Score []int
}

func (s Student) Average() float64 {
	var total int
	for _, e := range s.Score {
		total = total + e
	}
	avg := float64(total) / float64(len(s.Score))
	return avg

}

func (s Student) Min() (min int, name string) {
	for i, e := range s.Score {
		switch {
		case i == 0:
			min, name = e, s.Name[i]
		case e < min:
			min, name = e, s.Name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	for i, e := range s.Score {
		switch {
		case i == 0:
			max, name = e, s.Name[i]
		case e > max:
			max, name = e, s.Name[i]
		}
	}
	return max, name
}
