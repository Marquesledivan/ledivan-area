package area

import "math"

// Pi um PI
const Pi = 3.1455

// Circ csdhdsjkd
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi

}

// React csdhdsjkd
func React(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
