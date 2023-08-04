package shape

import (
	"math"
)

type Shape interface { // Интерфейс
	ShapeWithArea
	ShapeWithPerimeter // Создание метода интерфейса
}

type ShapeWithArea interface { // Интерфейс
	Area() float32 // Создание метода интерфейса
}

type ShapeWithPerimeter interface { // Интерфейс
	Perimeter() float32 // Создание метода интерфейса
}

type Square struct { // создание структуры
	sideLenght float32
}

func NewSquare(lenght float32) Square {
	return Square{
		sideLenght: lenght,
	}
}

func (s Square) Area() float32 { // определяем метод эриа
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimeter() float32 {
	return s.sideLenght * 4
}

type Circle struct { // создание структуры
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float32 {
	return c.radius
}
