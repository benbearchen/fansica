package formula

type Map interface {
	Value(x, y float64) float64
	X2Y(x float64) float64
}
