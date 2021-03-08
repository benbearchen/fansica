package formula

import "testing"

import (
	"math"
	"strings"
)

func near(a, b float64) bool {
	return math.Abs(a-b) < 1e-8
}

func geo(a, b, d float64) float64 {
	return math.Pow(a, 1-d) * math.Pow(b, d)
}

func TestCalcLine(t *testing.T) {
	check := func(xa, ya, xb, yb, x, v float64, e, e2 string) {
		line := Line{XY{xa, ya}, XY{xb, yb}, Liner}
		r, err := line.Value(x)
		r2, err2 := CalcLiner(xa, ya, xb, yb, x)

		if err != nil {
			if len(e) == 0 {
				t.Errorf("Line{%g, %g, %g, %g, Liner}.Value(%g) failed: %v", xa, ya, xb, yb, x, err)
			} else if strings.Index(err.Error(), e) < 0 {
				t.Errorf("Line{%g, %g, %g, %g, Liner}.Value(%g) unexpected failed: %v", xa, ya, xb, yb, x, err)
			}
		} else if len(e) > 0 {
			t.Errorf("Line{%g, %g, %g, %g, Liner}.Value(%g) don't fail, but want: %v", xa, ya, xb, yb, x, e)
		} else if !near(v, r) {
			t.Errorf("Line{%g, %g, %g, %g, Liner}.Value(%g) -> %g != %g", xa, ya, xb, yb, x, r, v)
		}

		if err2 != nil {
			if len(e2) == 0 {
				t.Errorf("CalcLiner(%g, %g, %g, %g, %g) failed: %v", xa, ya, xb, yb, x, err2)
			} else if strings.Index(err2.Error(), e2) < 0 {
				t.Errorf("CalcLiner(%g, %g, %g, %g, %g) unexpected failed: %v", xa, ya, xb, yb, x, err2)
			}
		} else if len(e2) > 0 {
			t.Errorf("CalcLiner(%g, %g, %g, %g, %g) don't fail, but want: %v", xa, ya, xb, yb, x, e)
		} else if !near(v, r2) {
			t.Errorf("CalcLiner(%g, %g, %g, %g, %g) -> %g != %g", xa, ya, xb, yb, x, r2, v)
		}

	}

	check(0, 1, 1, 2, 0, 1, "", "")
	check(0, 1, 1, 2, 1, 2, "", "")
	check(0, 1, 1, 2, 0.5, 1.5, "", "")
	check(0, 1, 1, 2, 0.1, 1.1, "", "")

	check(0, 1, 1, 2, -0.1, 0.9, "out of", "")

	check(0, 1, 0, 1, 0, 1, "", "")
	check(0, 1, 0, 1, 0.1, 0, "out of", "==")
	check(0, 1, 0, 1, -0.1, 0, "out of", "==")
}

func TestCalcReciprocal(t *testing.T) {
	check := func(xa, ya, xb, yb, x, v float64, e, e2 string) {
		line := Line{XY{xa, ya}, XY{xb, yb}, Reciprocal}
		r, err := line.Value(x)
		r2, err2 := CalcReciprocal(xa, ya, xb, yb, x)

		if err != nil {
			if len(e) == 0 {
				t.Errorf("Line{%g, %g, %g, %g, Reciprocal}.Value(%g) failed: %v", xa, ya, xb, yb, x, err)
			} else if strings.Index(err.Error(), e) < 0 {
				t.Errorf("Line{%g, %g, %g, %g, Reciprocal}.Value(%g) unexpected failed: %v", xa, ya, xb, yb, x, err)
			}
		} else if len(e) > 0 {
			t.Errorf("Line{%g, %g, %g, %g, Reciprocal}.Value(%g) don't fail, but want: %v", xa, ya, xb, yb, x, e)
		} else if !near(v, r) {
			t.Errorf("Line{%g, %g, %g, %g, Reciprocal}.Value(%g) -> %g != %g", xa, ya, xb, yb, x, r, v)
		}

		if err2 != nil {
			if len(e2) == 0 {
				t.Errorf("CalcReciprocal(%g, %g, %g, %g, %g) failed: %v", xa, ya, xb, yb, x, err2)
			} else if strings.Index(err2.Error(), e2) < 0 {
				t.Errorf("CalcReciprocal(%g, %g, %g, %g, %g) unexpected failed: %v", xa, ya, xb, yb, x, err2)
			}
		} else if len(e2) > 0 {
			t.Errorf("CalcReciprocal(%g, %g, %g, %g, %g) don't fail, but want: %v", xa, ya, xb, yb, x, e)
		} else if !near(v, r2) {
			t.Errorf("CalcReciprocal(%g, %g, %g, %g, %g) -> %g != %g", xa, ya, xb, yb, x, r2, v)
		}

	}

	check(0, 0, 0, 0, 0, 0, "<= 0", "<= 0")
	check(1, 0, 0, 0, 0, 0, "out of line", "<= 0")
	check(0, 0, 1, 0, 0, 0, "<= 0", "<= 0")
	check(1, 2, 0, 0, 0, 0, "out of line", "<= 0")
	check(1, 2, 0, 0, 0.1, 0, "out of line", "<= 0")
	check(1, 2, 2, 0, 1, 0, "<= 0", "<= 0")
	check(1, 2, 1, 2, 1, 2, "", "")
	check(1, 2, 1, 2, 1.1, 0, "out of line", "> B.X")

	check(1, 2, 2, 1, 1, 2, "", "")
	check(1, 2, 2, 1, 2, 1, "", "")
	check(1, 2, 2, 1, 1.1, 2/1.1, "", "")
	check(1, 2, 2, 1, 1.5, 2/1.5, "", "")
	check(1, 2, 2, 1, 1.8, 2/1.8, "", "")

	check(1, 3, 2, 1, 1, 3, "", "")
	check(1, 3, 2, 1, 2, 1, "", "")
	check(1, 3, 2, 1, 1.01, geo(3, 2, 0.01)/1.01, "", "")
	check(1, 3, 2, 1, 1.5, geo(3, 2, 0.5)/1.5, "", "")
	check(1, 3, 2, 1, 1.98, geo(3, 2, 0.98)/1.98, "", "")
}
