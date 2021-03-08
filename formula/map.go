package formula

import (
	"fmt"
)

type Map interface {
	Value(x, y float64) (float64, error)
	RangeOfX(x float64) (float64, float64, error)
}

func MakeFixMap(v float64, border Border) (Map, error) {
	vf := func(x, y float64) (float64, error) {
		return v, nil
	}

	return MakeFuncMap(vf, border)
}

func MakeFuncMap(v func(x, y float64) (float64, error), border Border) (Map, error) {
	return &simpleMap{v, border}, nil
}

type simpleMap struct {
	v func(x, y float64) (float64, error)
	b Border
}

func (m *simpleMap) Value(x, y float64) (float64, error) {
	min, max, err := m.b.RangeOfX(x)
	if err != nil {
		return 0, fmt.Errorf("out of border: %v", err)
	} else if y < min || y > max {
		return 0, fmt.Errorf("out of border, Y(%g) limit to <%g, %g>", y, min, max)
	} else {
		return m.v(x, y)
	}
}

func (m *simpleMap) RangeOfX(x float64) (float64, float64, error) {
	return m.b.RangeOfX(x)
}
