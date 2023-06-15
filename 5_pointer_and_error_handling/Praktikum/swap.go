package pointerAndErrorHandling

import "fmt"

func Swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
	fmt.Println(*a, *b)
}
