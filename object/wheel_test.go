package object

import "testing"

import (
	"github.com/benbearchen/fansica/formula"
)

func TestParseWheel(t *testing.T) {
	check := func(spec string, d formula.Millimeter) {
		w, err := ParseWheel(spec)
		if err != nil {
			if d != 0 {
				t.Errorf("ParseWheel(%s) failed: %v", spec, err)
			}

			return
		}

		if d == 0 {
			t.Errorf("ParseWheel(%s) should fail, but return: %v", spec, w)
			return
		}

		p := formula.ToMillimeter(w.Perimeter())
		if int(p) != int(d) {
			t.Errorf("ParseWheel(%s).Perimeter() -> %v != %v", spec, p, d)
		}
	}

	check("205/50R15", 1840)
	check("205/55R16", 1985)
	check("195/65R15", 1993)

	check("195/65R15V", 1993)
	check("195/65r15", 0)
	check("195/R15V", 0)
	check("/65R15V", 0)
	check("R/65R15V", 0)
	check("R195/65R15V", 0)
}
