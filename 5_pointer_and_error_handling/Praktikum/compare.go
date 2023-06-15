package pointerAndErrorHandling

import "strings"

func CompareString(a, b string) string {
	if res := strings.Contains(b, a); res {
		return a
	}
	return b

}
