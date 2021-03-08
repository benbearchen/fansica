package formula

import (
	"fmt"
	"math"
)

type LineType int

const (
	Liner      LineType = iota // 直线
	Reciprocal                 // 倒数
)

type XY struct {
	X float64
	Y float64
}

type Line struct {
	A    XY
	B    XY
	Type LineType
}

type Border interface {
	RangeOfX(x float64) (float64, float64, error)
}

type basicBorder struct {
	baseY float64
	lines []Line
}

func MakeBorder(baseY float64, p ...interface{}) (Border, error) {
	lines := make([]Line, 0, len(p))
	var next *Line
	for _, v := range p {
		switch v := v.(type) {
		case LineType:
			if next != nil {
				next.Type = v
			}
		case XY:
			if next != nil {
				if next.A.X >= v.X {
					return nil, fmt.Errorf("prev X(%g) > next X(%g)", next.A.X, v.X)
				}

				next.B = v
				lines = append(lines, *next)
				next = nil
			}

			next = new(Line)
			next.A = v
		default:
			return nil, fmt.Errorf("unknown type: %v", v)
		}
	}

	if len(lines) < 1 {
		return nil, fmt.Errorf("can't make lines from: %v", p)
	}

	return &basicBorder{baseY, lines}, nil
}

func (line *Line) Has(x float64) bool {
	return x >= line.A.X && x <= line.B.Y
}

func (line *Line) Value(x float64) (float64, error) {
	if x < line.A.X || x > line.B.X {
		return 0, fmt.Errorf("X(%g) out of line.X(%g <-> %g)", x, line.A.X, line.B.X)
	}

	switch line.Type {
	case Liner:
		return CalcLiner(line.A.X, line.A.Y, line.B.X, line.B.Y, x)
	case Reciprocal:
		return CalcReciprocal(line.A.X, line.A.Y, line.B.X, line.B.Y, x)
	default:
		return 0, fmt.Errorf("wrong line type: %v", line.Type)
	}
}

func CalcLiner(xa, ya, xb, yb, x float64) (float64, error) {
	if xa == x {
		return ya, nil
	} else if xb == x {
		return yb, nil
	}

	dx := xb - xa
	if dx == 0 {
		return 0, fmt.Errorf("A.X(%g) == B.X(%g)", xa, xb)
	}

	y := ya + (yb-ya)*(x-xa)/dx
	return y, nil
}

func CalcReciprocal(xa, ya, xb, yb, x float64) (float64, error) {
	if xa <= 0 {
		return 0, fmt.Errorf("A.X(%g) <= 0", xa)
	}

	if xb <= 0 {
		return 0, fmt.Errorf("B.X(%g) <= 0", xb)
	}

	if x <= 0 {
		return 0, fmt.Errorf("X(%g) <= 0", x)
	}

	if ya <= 0 {
		return 0, fmt.Errorf("A.Y(%g) <= 0", ya)
	}

	if yb <= 0 {
		return 0, fmt.Errorf("B.Y(%g) <= 0", yb)
	}

	if x < xa {
		return 0, fmt.Errorf("X(%g) < A.X(%g)", x, xa)
	}

	if x > xb {
		return 0, fmt.Errorf("X(%g) > B.X(%g)", x, xb)
	}

	if xa > ya {
		return 0, fmt.Errorf("A.X(%g) > B.X(%g)", xa, xb)
	}

	if x == xa {
		return ya, nil
	} else if x == xb {
		return yb, nil
	}

	dx := xb - xa
	if dx == 0 {
		return 0, fmt.Errorf("A.X(%g) == B.X(%g)", xa, xb)
	}

	sa := xa * ya
	sb := xb * yb

	s := math.Pow(sa, (xb-x)/dx) * math.Pow(sb, (x-xa)/dx)
	return s / x, nil
}

func (bb *basicBorder) RangeOfX(x float64) (float64, float64, error) {
	for _, line := range bb.lines {
		if line.Has(x) {
			v, err := line.Value(x)
			return bb.baseY, v, err
		}
	}

	return bb.baseY, 0, fmt.Errorf("%g out of border", x)
}
