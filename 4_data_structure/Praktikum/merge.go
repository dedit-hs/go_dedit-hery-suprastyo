package dataStructure

func Merge(arrayA, arrayB []string) []string {
	var uniqueElements = make(map[string]bool)
	for _, el := range arrayA {
		uniqueElements[el] = true
	}

	for _, el := range arrayB {
		uniqueElements[el] = true
	}

	result := make([]string, 0, len(uniqueElements))
	for element := range uniqueElements {
		result = append(result, element)
	}

	return result
}
