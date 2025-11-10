package main

import "fmt"

func sum(arr [10]int) int {

	s := 0

	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}

	return s
}

func multiply(s []int, factor int, doCopy bool) []int {

	if doCopy {
		sCopy := make([]int, len(s), len(s))
		copy(sCopy, s)
		s = sCopy
	}

	for i := 0; i < len(s); i++ {
		s[i] = s[i] * factor
	}

	return s
}

func mainDiagonal(m [][]int) []int {

	next := 0
	var result []int

	for i := 0; i < len(m); i++ {
		result = append(result, m[i][next])
		next++
	}

	return result
}

func secondaryDiagonal(m [][]int) []int {

	next := len(m) - 1
	var result []int

	for i := 0; i < len(m); i++ {
		result = append(result, m[i][next])
		next--
	}

	return result
}

func main() {
	// Задание 1. Сумма элементов массива
	fmt.Printf("sum is %d\n", sum([...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// sum is 55

	// Задание 2. Умножение всех элементов слайса
	sl1 := []int{1, 2, 3, 4, 5}
	res1 := multiply(sl1, 3, false)
	fmt.Println("Полсе выполнения multiply() без копирования sl1 стал ", sl1, "; резултат = ", res1)
	// Полсе выполнения multiply() без копирования sl1 стал  [3 6 9 12 15] ; резултат =  [3 6 9 12 15]

	sl2 := []int{1, 2, 3, 4, 5}
	res2 := multiply(sl2, 3, true)
	fmt.Println("Полсе выполнения multiply() с копированием sl2 не изменился ", sl2, "; резултат = ", res2)
	// Полсе выполнения multiply() с копированием sl2 не изменился  [1 2 3 4 5] ; резултат =  [3 6 9 12 15]

	// Задание 3. Обрезка и копирование
	numbers := make([]int, 0, 20)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Адрес numbers[3] %p len(%d) cap(%d)\n", &numbers[3], len(numbers), cap(numbers))
	// Адрес numbers[3] 0xc0000200b8 len(10) cap(20)

	part := numbers[3:8]
	fmt.Printf("Адрес part[0] %p len(%d) cap(%d)\n", &part[0], len(part), cap(part))
	// Адрес part[0] 0xc0000200b8 len(5) cap(17) - такой же как у 3 элемента numbers!

	copyPart := make([]int, len(part))
	copy(copyPart, part)
	fmt.Printf("Адрес copyPart[0] %p len(%d) cap(%d)\n", &copyPart[0], len(copyPart), cap(copyPart))
	// Адрес copyPart[0] 0xc00012e090 len(5) cap(5) - адрес изменился!

	numbers[3] = 50
	fmt.Println("numbers:", numbers, "; part:", part, "; copyPart:", copyPart)
	// numbers: [1 2 3 50 5 6 7 8 9 10] ; part: [50 5 6 7 8] ; copyPart: [4 5 6 7 8]

	// Задание 4. Матрица 3×3 и диагонали
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, len(matrix))
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
		fmt.Println(matrix[i])
	}
	/*
		[0 1 2]
		[1 2 3]
		[2 3 4]
	*/
	fmt.Println("mainDiagonal() result is:", mainDiagonal(matrix))
	// mainDiagonal() result is: [0 2 4]

	fmt.Println("secondaryDiagonal() result is:", secondaryDiagonal(matrix))
	// secondaryDiagonal() result is: [2 2 2]
}
