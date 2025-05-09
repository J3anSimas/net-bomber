package vector

import (
	"math"
	"math/rand/v2"
)

type Point struct {
	x, y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{x, y}
}

func (p *Point) X() float32 {
	return p.x
}
func (p *Point) Y() float32 {
	return p.y
}
func (p *Point) Add(p2 *Point) {
	p.x = float32(math.Round(float64(p.x+p2.x)*100000) / 100000)
	p.y = float32(math.Round(float64(p.y+p2.y)*100000) / 100000)
}
func (p *Point) Sub(p2 *Point) {
	p = &Point{p.x - p2.x, p.y - p2.y}
}
func (p *Point) Mult(p2 *Point) {
	p = &Point{p.x * p2.x, p.y * p2.y}
}
func (p *Point) Div(p2 *Point) {
	p = &Point{p.x / p2.x, p.y / p2.y}
}
func (p *Point) MultScalar(scalar float32) {
	p = &Point{p.x * scalar, p.y * scalar}
}
func (p *Point) DivScalar(scalar float32) {
	p = &Point{p.x / scalar, p.y / scalar}
}
func (p *Point) Abs() {
	p = &Point{
		x: float32(math.Abs(float64(p.x))),
		y: float32(math.Abs(float64(p.y))),
	}
}
func (p *Point) LimitScalar(limit float32) {
	if p.x > limit {
		p.x = limit
	}
	if p.x < -limit {
		p.x = -limit
	}
	if p.y > limit {
		p.y = limit
	}
	if p.y < -limit {
		p.y = -limit
	}
}
func Up() *Point {
	return &Point{0, 1}
}
func Down() *Point {
	return &Point{0, -1}
}
func Left() *Point {
	return &Point{-1, 0}
}
func Right() *Point {
	return &Point{1, 0}
}
func Zero() *Point {
	return &Point{0, 0}
}
func Random() *Point {

	x := float32(rand.Float32())
	polarity := rand.Float32()
	if polarity < 0.5 {
		x *= -1
	}
	y := float32(rand.Float32())
	polarity = rand.Float32()
	if polarity < 0.5 {
		y *= -1
	}
	return &Point{
		x: x,
		y: y,
	}
}
