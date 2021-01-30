package formula

import "testing"

import (
	"log"
)

func TestInterpolationWithBSFC(t *testing.T) {
	kw := []float64{
		7.1,
		9.7,
		26.8,
		15.5,
	}

	bsfc := []float64{
		0.35,
		0.38,
		0.43,
		0.40,
	}

	inter := NewInterpolation(kw, bsfc)
	kw0 := Kilowatt(7)
	fuel := float64(kw0) / bsfc[0] / 9
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
		log.Printf("%v*%g*1h <- %.4gL -> %v*%.3g*%.3gh, trans: %.4gkWh, total:%.4gkWh, %.2f%%", kw0, bsfc[0], fuel, k, c, t, e1, ee, (ee/float64(kw0)-1)*100)
	}
}

func TestInterpolationWithBSFC2(t *testing.T) {
	kw := []float64{
		5.45,
		8,
		20,
		31.4,
	}

	bsfc := []float64{
		0.32,
		0.33,
		0.34,
		0.347,
	}

	inter := NewInterpolation(kw, bsfc)
	kw0 := Kilowatt(6)
	fuel := float64(kw0) / bsfc[0] / 9
	for i := 0; i < 50; i++ {
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
		e1 := (e - float64(kw0)*t) * 0.90
		ee := float64(kw0)*t + e1
		log.Printf("%v*%g*1h <- %.4gL -> %v*%.3g*%.3gh, trans: %.4gkWh, total:%.4gkWh, %.2f%%", kw0, bsfc[0], fuel, k, c, t, e1, ee, (ee/float64(kw0)-1)*100)
	}
}
