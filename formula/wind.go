package formula

const (
	NormalAirDensity = 1.205
)

func CalcAirResistance(cd, density, v, s float64) float64 {
	return cd * density * s * v * v / 2
}
