package formula

import "testing"

import (
	"math"
)

func TestCalcRotatorPower(t *testing.T) {
	check := func(rpm, nm, kw float64) {
		r := CalcRotatorPower(RotationPerMinute(rpm), NewtonMeter(nm))
		d := math.Abs(float64(r)/kw - 1)
		if d > 1e-3 {
			t.Errorf("CalcRotatorPower(%v, %v) -> %v != %v", RotationPerMinute(rpm), NewtonMeter(nm), r, Kilowatt(kw))
		}
	}

	check(954.93, 20, 2)
}
