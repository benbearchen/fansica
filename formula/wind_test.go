package formula

import "testing"

import (
	"log"
)

func TestCalcAirResistance(t *testing.T) {
	width := 1.96
	height := 1.76

	s := width * height
	cd := 0.32

	var v KilometerPerHour = 120

	f := CalcAirResistance(cd, NormalAirDensity, SI(v), s)
	if int(f) != 738 {
		t.Errorf("CalcAirResistance(%.3g, %.4gkg/m3, %v, %.5gm2) -> %.2fN != 738N", cd, NormalAirDensity, v, s, f)
	} else {
		power := ToKilowatt(SI(f) * SI(v))
		log.Printf("CalcAirResistance(%.3g, %.4gkg/m3, %v, %.5gm2) -> %.2fN (%v)", cd, NormalAirDensity, v, s, f, power)
	}
}
