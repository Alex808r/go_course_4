package interview

import (
	"fmt"
	"sync"
)

// TypicalErrorsNoSync демонстрирует две классические проблемы:
// 1. Отсутствие синхронизации (горутины запускаются, но main завершается до их выполнения).
// 2. До Go 1.22 переменная цикла `i` захватывалась по ссылке (все горутины печатали последнее значение, обычно 10).
func TypicalErrorsNoSync() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

// TypicalErrorsIndex решает проблему завершения с помощью sync.WaitGroup.
// Идиоматично передавать значение итератора как аргумент горутины: `go func(n int) { ... }(i)`.
func TypicalErrorsIndex() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
}

// Slices демонстрирует механику срезов в Go:
// 1. Срез s2 изначально указывает на тот же базовый массив, что и s1.
//    Поэтому s2[1] = 999 изменяет и s1[2].
// 2. При append, если превышена вместимость (cap) среза, Go выделяет НОВЫЙ базовый массив.
//    После реалокации s1 и s2 ссылаются на разные участки памяти.
func Slices() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[1:3]
	fmt.Println("До изменения S2")
	fmt.Println(s1)
	fmt.Println(s2)

	s2[1] = 999
	fmt.Println("После изменения S2 (s1 тоже изменился, т.к. общий массив):")
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 5, 6, 7, 8, 9)
	fmt.Println("После добавления к S2 (произошла реалокация нового массива):")
	fmt.Println(s1)
	fmt.Println(s2)
}

// Strings демонстрирует итерацию по UTF-8 строке:
// `range s` декодирует руны (UTF-8 символы), возвращая индекс байта начала руны и код руны (int32).
func Strings() {
	s := "Привет!"

	for index, runeVal := range s {
		fmt.Printf("байт: %d, руна: %c (код: %d)\n", index, runeVal, runeVal)
	}
}
