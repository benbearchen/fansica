package formula

import "testing"

import (
	"log"
)

func TestInterpolationWithTE(t *testing.T) {
	kw := []float64{
		7.1,
		9.7,
		15.5,
		26.8,
	}

	te := []float64{
		0.35,
		0.38,
		0.40,
		0.43,
	}

	inter := NewInterpolation(kw, te)
	kw0 := Kilowatt(7)
	te0, _ := inter.Value(float64(kw0))
	fuel := float64(kw0) / te0 / 9
	tops := make(map[float64]float64)
	for i := 0; i < 41; i++ {
		k := kw0 + Kilowatt(float64(i)/2)
		c, ok := inter.Value(float64(k))
		if !ok {
			log.Printf("Interpolation[%v] out of range", k)
			continue
		} else {
			log.Printf("Interpolation[%v] = %g", k, c)
		}

		e := fuel * 9 * c
		t := e / float64(k)
		e1 := (e - float64(kw0)*t) * 0.95
		ee := float64(kw0)*t + e1
		log.Printf("%v*%g*1h <- %.4gL -> %v*%.3g*%.3gh, trans: %.4gkWh, total:%.4gkWh, %.2f%%", kw0, te[0], fuel, k, c, t, e1, ee, (ee/float64(kw0)-1)*100)

		cTop, kwTop := te[len(te)-1], kw[len(kw)-1]
		per := CalcThermalEfficiencyAdvance(c, float64(k), cTop, kwTop, 0.95)
		tops[float64(k)] = per
	}

	log.Printf("promote to 43%%: %v", tops)
}

func TestInterpolationWithTE2(t *testing.T) {
	kw := []float64{
		2.89,
		3.14,
		3.54,
		4.0,
		4.6,
		5.4,
		8,
		20,
		31.4,
	}

	te := []float64{
		0.27,
		0.28,
		0.29,
		0.30,
		0.31,
		0.32,
		0.33,
		0.34,
		0.347,
	}

	inter := NewInterpolation(kw, te)
	kw0 := Kilowatt(3)
	te0, _ := inter.Value(float64(kw0))
	fuel := float64(kw0) / te0 / 9
	lows := make(map[float64]float64)
	tops := make(map[float64]float64)
	for i := 0; i < 60; i++ {
		k := kw0 + Kilowatt(float64(i)/2)
		c, ok := inter.Value(float64(k))
		if !ok {
			log.Printf("Interpolation[%v] out of range", k)
			continue
		} else {
			log.Printf("Interpolation[%v] = %g", k, c)
		}

		e := fuel * 9 * c
		t := e / float64(k)
		charge := 0.95
		e1 := (e - float64(kw0)*t) * charge
		ee := float64(kw0)*t + e1
		log.Printf("%v*%g*1h <- %.4gL -> %v*%.3g*%.3gh, trans: %.4gkWh, total:%.4gkWh, %.2f%%", kw0, te[0], fuel, k, c, t, e1, ee, (ee/float64(kw0)-1)*100)
		lows[float64(k)] = ee / float64(kw0)

		cTop, kwTop := te[len(te)-1], kw[len(kw)-1]
		per := CalcThermalEfficiencyAdvance(c, float64(k), cTop, kwTop, charge)
		log.Printf("%v*%.3g / %v*%.3g*%.3g => %.2f%%", k, c, Kilowatt(kwTop), cTop, charge, (per-1)*100)
		tops[float64(k)] = per
	}

	log.Printf("%v", lows)
	log.Printf("%v", tops)
}
