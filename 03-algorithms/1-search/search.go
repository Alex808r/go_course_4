package bsearch

import "fmt"

// Simple возвращает индекс элемента в слайсе или -1. Используется линейный поиск.
// Работает на любых (в том числе неотсортированных) слайсах.
// Временная сложность: O(N), дополнительная память: O(1).
func Simple(data []int, item int) int {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}

// Binary возвращает индекс элемента в слайсе или -1. Используется бинарный поиск.
// ВАЖНО: работает ТОЛЬКО для отсортированных слайсов (в данном случае — по возрастанию).
// Если слайс не отсортирован, алгоритм даст некорректный результат (-1), так как
// опирается на предположение, что элементы упорядочены, и отбрасывает половину данных.
// Временная сложность: O(log N), дополнительная память: O(1).
// Примечание: вычисление mid := low + (high-low)/2 предотвращает целочисленное
// переполнение при больших массивах (в отличие от (low + high) / 2).
func Binary(data []int, item int) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == item {
			return mid
		}
		if data[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func O() {
	slice := []int{1, 2, 3}

	// O(1)
	fmt.Println(len(slice))

	// O(N)
	for i := range slice {
		// ....

		// O(N^2)
		for j := range slice {

			// O(N^3)
			for k := range slice {
				_, _, _ = i, j, k
			}
		}
	}
}
