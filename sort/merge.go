package sort

// Implementação do algoritmo Merge Sort
// Passos do algoritmo:
// 1. Dividir o array em duas metades até que cada subarray tenha um único elemento.
// 2. Combinar os subarrays ordenados em um único array ordenado.
// 3. Repetir o processo até que todo o array esteja ordenado.
// O algoritmo tem complexidade de tempo O(n log n) e complexidade de espaço O(n).

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	mid := len(input) / 2
	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}
	return result
}
