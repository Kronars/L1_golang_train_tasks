package main

import "math"
import "fmt"

// Структура с инкапсулируемыми полями
type point struct {
	x, y float64 // Будут недоступны из внешних пакетов так как с маленькой буквы
}

// Получение координат
func (p point) Coords() (float64, float64) {
	return p.x, p.y
}

// Метод нахождения расстояния
func (p point) Distance(p2 PointFace) float64 {
	x2, y2 := p2.Coords()
	dx := p.x - x2
	dy := p.y - y2
	return math.Sqrt(dx*dx + dy*dy)
}

// Интерфейс для инкапсуляции
type PointFace interface {
	Coords() (float64, float64)
	Distance(p2 PointFace) float64
}

// Встраивание интерфейса
type Point struct {
	PointFace
}

// Констрктор
func NewPoint(x, y float64) Point {
	return Point{point{x: x, y: y}}
}

func main() {
	p := NewPoint(5, 5)
	q := NewPoint(10, 10)
	fmt.Printf("Тип: %#v\n\n", p)

	d := q.Distance(p)
	fmt.Println(d)
}
