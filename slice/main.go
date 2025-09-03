package main

import (
	"fmt"
	"slices"
)

func main() {
	// Clip удаляет неиспользуемую память из фрагмента. Это может быть использовано для освобождения памяти обратно во время выполнения,
	// если в фрагменте было большое количество элементов, но с тех пор его длина сократилась (или если элементы используются).
	ints := make([]int, 4, 6)

	ints[0] = 1
	ints[1] = 2
	fmt.Printf("len: %d, capacity: %d\n", len(ints), cap(ints)) // len: 4, capacity: 6
	ints = slices.Clip(ints)
	fmt.Printf("len: %d, capacity: %d\n", len(ints), cap(ints)) // len: 4, capacity: 4

	// Clone возвращает копию фрагмента. Все значения копируются путем присваивания,
	// что означает, что это считается неполной копией элементов фрагмента. Неполная копия означает,
	// что исходные адреса памяти все еще используются, поэтому возможно, что вы все еще можете вносить
	// изменения в исходные данные. Хотя обычно это не является проблемой, важно иметь это в виду,
	// если вы хотите получить действительно независимую копию данных.
	original := []int{0, 1, 2, 4, 5}
	cloned := slices.Clone(original)
	cloned[0] = 9
	cloned[2] = 7
	cloned[3] = 6
	fmt.Printf("original: %v\n", original) // original: [0 1 2 4 5]
	fmt.Printf("cloned  : %v\n", cloned)   // cloned  : [9 1 7 6 5]

	// Contains - Сообщает, существует ли указанное значение в срезе.
	s := []string{"Apple", "Banana", "Orange"}
	fmt.Println("Found Apple?", slices.Contains(s, "Apple"))           // Found Apple? true
	fmt.Println("Found Banana?", slices.Contains(s, "Banana"))         // Found Banana? true
	fmt.Println("Found Strawberry?", slices.Contains(s, "Strawberry")) // Found Strawberry? false

	//Функция Index возвращает индекс первого вхождения искомого значения в срез. Возвращает -1, если значение отсутствует
	fmt.Println("Index for Apple:      ", slices.Index(s, "Apple"))      // Index for Apple: 0
	fmt.Println("Index for Banana:     ", slices.Index(s, "Banana"))     // Index for Banana: 1
	fmt.Println("Index for Strawberry: ", slices.Index(s, "Strawberry")) // Index for Strawberry: -1

	//Insert вставляет значение по указанному индексу. Срез будет увеличиваться по мере необходимости, освобождая место для вставленных значений.
	s = slices.Insert(s, 2, "Grape")
	fmt.Println(s) // [Apple Banana Grape Orange]

	// Max, Min
	ints2 := []int{1, 2, 10, 228, 14, 111}
	fmt.Println("Max value for s slice is", slices.Max(s))         // Max value for s slice is Orange
	fmt.Println("Max value for ints2 slice is", slices.Max(ints2)) // Max value for ints2 slice is 228

	fmt.Println("Min value for s slice is", slices.Min(s))         // Min value for s slice is Apple
	fmt.Println("Min value for ints2 slice is", slices.Min(ints2)) // Min value for ints2 slice is 1

	// Replace - заменяет выбранную часть среза указанными значениями.

	s2 := []string{"Maxim", "Anatoly", "Masha", "Sasha"}

	s2 = slices.Replace(s2, 1, 2, "Kolya", "Peter")

	fmt.Println(s2) // [Maxim Kolya Peter Masha Sasha]

	// Sort - сортировка

	ints3 := []int{1, 14535, 10, 228, 14, 111}

	slices.Sort(ints3)

	fmt.Println(ints3) // [1 10 14 111 228 14535]
}
