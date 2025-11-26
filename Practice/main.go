package main

import (
	"fmt"
	"strings"
)

func CountInts(nums []int) map[int]int {

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	return m
}

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	// Убираем запятые и приводим к нижнему регистру
	words := strings.Fields(strings.ToLower(strings.Replace(s, ",", "", -1)))
	for _, word := range words {
		m[word]++
	}

	return m
}

func IsAnagram(a, b string) bool {

	inv := []rune(a)
	for i, j := 0, len(inv)-1; i < j; i, j = i+1, j-1 {
		inv[i], inv[j] = inv[j], inv[i]
	}

	return strings.ToLower(string(inv)) == strings.ToLower(b)
}

func Invert(m map[string]int) map[int]string {

	inverted := make(map[int]string)

	for k, v := range m {
		inverted[v] = k
	}

	return inverted
}

func MergeMaps(dst, src map[string]int) map[string]int {

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

type User struct {
	Name string
	City string
}

func newUser(name string, city string) User {
	return User{
		Name: name,
		City: city,
	}
}

func GroupByCity(users []User) map[string][]User {
	g := make(map[string][]User)

	for _, user := range users {
		g[user.City] = append(g[user.City], user)
	}

	return g
}

func main() {

	// 1 почитать количество одинаковых цифр в слайсе
	i := []int{2, 3, 4, 5, 3, 2, 2, 4, 5, 6, 7, 5, 3, 2, 2, 4, 5}
	fmt.Println("Количестов одинаковых цифр в", i, "Результат:", CountInts(i))

	// 2 Посчитать количество слов в строке
	s := "Эх, если бы да кабы, во рту грибы росли, то был бы не рот, а целый огород"
	fmt.Println(fmt.Sprintf("Количество слов в строке: \"%s\"\nРезультат:", s), WordCount(s))

	// 3 Проверить, является ли слово Анаграммой
	s1 := "Лана"
	s2 := "анал"
	var answer string
	if IsAnagram(s1, s2) {
		answer = "да является )"
	} else {
		answer = "к сожалению нет"
	}

	fmt.Printf("Является ли \"%s\" анаграмой слова \"%s\" ответ: %s\n", s1, s2, answer)

	// 4 Сделать инверсию мапы
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("Мапа до инверсии:", m, "\nМапа после инверсии:", Invert(m))

	// 5 слияние 2х мап
	m1 := map[string]int{
		"a": 3,
		"c": 5,
		"f": 3,
	}
	m2 := map[string]int{
		"f": 2,
		"d": 4,
		"a": 5,
	}
	fmt.Println("Слияние мапы\n", m1, "\nc мапой\n", m2)
	fmt.Println("Результат:", MergeMaps(m1, m2))

	// 6 Сгруппировать всех юзеров по городу
	u1 := newUser("Федя", "Алматы")
	u2 := newUser("Кирилл", "Москва")
	u3 := newUser("Алибек", "Чу")
	u4 := newUser("Володя", "Алматы")
	u5 := newUser("Дима", "Алматы")
	users := []User{u1, u2, u3, u4, u5}
	fmt.Println("Группируем юзеров:\n", users, "\nпо городам\nРезульат:\n", GroupByCity(users))
}
