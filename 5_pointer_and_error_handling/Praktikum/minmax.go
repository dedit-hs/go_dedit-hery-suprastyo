package pointerAndErrorHandling

func GetMinMax(number ...*int) (min, max int) {
	for i, e := range number {
		switch {
		case i == 0:
			min, max = *e, *e
		case *e > max:
			max = *e
		case *e < min:
			min = *e
		}
	}
	return min, max
}
