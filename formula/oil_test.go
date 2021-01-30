package formula

import "testing"

func TestCalcGasoline100km(t *testing.T) {
	kph := KilometerPerHour(120)
	bsfc := 0.33
	kw := Kilowatt(36)

	lphkm := ToLiterPer100Kilometer(CalcGasoline100km(SI(kw), bsfc, SI(kph)))
	if int(lphkm) != 10 {
		t.Errorf("CalcGasoline100km(%v, %.3g%%, %v) -> %v", kw, bsfc*100, kph, lphkm)
	}
}
