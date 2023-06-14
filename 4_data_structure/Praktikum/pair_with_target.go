package dataStructure

func PairSum(arr []int, target int) []int {
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {

			if expectedTarget := arr[i] + arr[j]; expectedTarget == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
