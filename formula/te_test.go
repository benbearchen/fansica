package formula

import "testing"

import (
	"math"
)

func TestCalcThermalEfficiencyAdvance(t *testing.T) {
	check := func(clow, wlow, chi, whi, charge, a float64) {
		r := CalcThermalEfficiencyAdvance(clow, wlow, chi, whi, charge)
		if math.Abs(r-a) > 1e-8 {
			t.Errorf("CalcThermalEfficiencyAdvance(%g%%, %g, %g%%, %g, %g%%) -> %g != %g", clow, wlow, chi, whi, charge, r, a)
		}
	}

	check(0.20, 1, 0.33, 8, 0.9, 1.505625)
}
