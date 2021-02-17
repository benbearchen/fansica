package object

import "testing"

func TestNamer(t *testing.T) {
	check := func(n Namer, obj Object, a, b, c string, tag string) {
		if n.Name() != a {
			t.Errorf("%s.Name() -> %s != %s", tag, n.Name(), a)
		}

		if n.TypeName(obj) != b {
			t.Errorf("%s.Name() -> %s != %s", tag, n.TypeName(obj), b)
		}

		if n.NamedString(obj) != c {
			t.Errorf("%s.Name() -> %s != %s", tag, n.NamedString(obj), c)
		}
	}

	wheel, _ := ParseWheel("wheel0", "255/50R20")
	check(wheel, wheel, "wheel0", "Wheel", "wheel0(Wheel)", "Wheel")

	eng := EasyMakeEngine("eng0", 6500, 220, 96)
	check(eng, eng, "eng0", "Engine", "eng0(Engine)", "Engine")

	b := EasyMakeBattery("bat", 40, 220)
	check(b, b, "bat", "Battery", "bat(Battery)", "Battery")

	m := EasyMakeMotor("mg", 100, 240, 12000)
	check(m, m, "mg", "Motor", "mg(Motor)", "Motor")

	pcu := NewPCU("control", 4)
	check(pcu, pcu, "control", "PCU", "control(PCU)", "PCU")
}
