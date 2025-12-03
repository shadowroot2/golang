package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) DistanceFromOrigin() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

type Robot struct {
	ID string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет я %s\n", p.Name)
}

func (r Robot) Greet() string {
	return fmt.Sprintf("Привет кожанные, я %s\n", r.ID)
}

type Updater interface {
	Update(msg string)
}

type Logger struct {
	LastMessage string
}

func (l *Logger) Update(msg string) {
	l.LastMessage = msg
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	// Задача 1. Простейший метод со значением-получателем
	p := Point{X: -4, Y: 8}
	fmt.Println("Дистанция:", p.DistanceFromOrigin())

	// Задача 2. Метод с указателем-получателем (изменение состояния)
	p.Move(3, -1)
	fmt.Println("Точка:", p.X, p.Y) // 4 1

	// Задача 3 - делали на практике уже

	// Задача 4. Полиморфизм: несколько реализаций одного интерфейса
	p1 := Person{"Костя"}
	r1 := Robot{"ZBL-21"}
	SayHello(p1)
	SayHello(r1)

	// Задача 5. Срез интерфейсов
	var greeters []Greeter
	greeters = append(greeters, Person{"Вова"}, Person{"Вася"}, Person{"Кирилл"}, Robot{"Пылесос"}, Robot{"Хуесос"})
	for _, g := range greeters {
		fmt.Println(g.Greet())
	}

	// Задача 6. Указатель/значение и интерфейс
	// var u1 Updater
	var u2 Updater

	l := Logger{}

	// u1 = l      // <- так не компилируется
	u2 = &l // <- а так компилируется

	u2.Update("Привед медвед")
	fmt.Println("Последняя запись лога:", l.LastMessage)
}
