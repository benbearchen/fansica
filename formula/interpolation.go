package formula

import (
	"sort"
)

type Interpolation struct {
	x []float64
	y []float64
}

func NewInterpolation(x, y []float64) *Interpolation {
	if len(x) <= 0 || len(y) <= 0 {
		return new(Interpolation)
	}

	a := make([]int, len(x))
	if len(a) > len(y) {
		a = a[:len(y)]
	}

	for i := range a {
		a[i] = i
	}

	sort.Slice(a, func(i, j int) bool {
		return x[a[i]] < x[a[j]]
	})

	xa := make([]float64, len(a))
	ya := make([]float64, len(a))
	for i, k := range a {
		xa[i] = x[k]
		ya[i] = y[k]
	}

	return &Interpolation{xa, ya}
}

func (inter *Interpolation) Value(x float64) (float64, bool) {
	s := len(inter.x)
	if s <= 0 {
		return 0, false
	}

	d := x - inter.x[0]
	if d < 0 {
		return 0, false
	} else if d == 0 {
		return inter.y[0], true
	}

	d = x - inter.x[s-1]
	if d > 0 {
		return 0, false
	} else if d == 0 {
		return inter.y[s-1], true
	}

	p := sort.SearchFloat64s(inter.x, x)

	x0, y0 := inter.x[p-1], inter.y[p-1]
	x1, y1 := inter.x[p], inter.y[p]

	y := y0 + (y1-y0)*(x-x0)/(x1-x0)
	return y, true
}
