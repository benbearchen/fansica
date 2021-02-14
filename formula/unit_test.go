package formula

import "testing"

import (
	"log"
)

func TestUnit(t *testing.T) {
	var kw Kilowatt = 240
	if ToKilowatt(SI(kw)) != kw {
		t.Errorf("ToKilowatt(SI(%v)) -> %v", kw, ToKilowatt(SI(kw)))
	}

	if ToKilowatt(kw) != kw {
		t.Errorf("ToKilowatt(%v) -> %v", kw, ToKilowatt(kw))
	}

	ps := ToPS(SI(kw))
	if int(ps) != 326 {
		t.Errorf("%v to %v, not 326PS", kw, ps)
	} else {
		log.Printf("%v == %v", kw, ps)
	}

	if ToPS(SI(ps)) != ps {
		t.Errorf("ToPS(SI(%v)) -> %v", ps, ToPS(SI(ps)))
	}

	ps = 306
	hp := ToHorsePower(SI(ps))
	if int(hp) != 301 {
		t.Errorf("%v to %v, not 301HP", ps, hp)
	} else {
		log.Printf("%v == %v", ps, hp)
	}

	if ToHorsePower(SI(hp)) != hp {
		t.Errorf("ToHorsePower(SI(%v)) -> %v", hp, ToHorsePower(SI(hp)))
	}

	kWh := KilowattHour(9)
	mj := SI(kWh) / 0.725 / 1000 / 1000
	if int(mj) != 44 {
		t.Errorf("9kWh/L to %.4gMJ/kg, not 44MJ/kg", mj)
	} else {
		log.Printf("9kWh/L to %.4gMJ/kg", mj)
	}

	if ToKilowattHour(SI(kWh)) != kWh {
		t.Errorf("ToKilowattHour(SI(%v)) -> %v", kWh, ToKilowattHour(SI(kWh)))
	}

	lphkm := LiterPer100Kilometer(1)
	if ToLiterPer100Kilometer(SI(lphkm)) != lphkm {
		t.Errorf("ToLiterPer100Kilometer(SI(%v)) -> %v", lphkm, ToLiterPer100Kilometer(SI(lphkm)))
	}

	rpm := RotationPerMinute(3000)
	if rpm.SI() != 50 {
		t.Errorf("RotationPerMinute(%v).SI() -> %f != 50", rpm, rpm.SI())
	}

	if ToRotationPerMinute(SI(rpm)) != rpm {
		t.Errorf("ToRotationPerMinute(SI(%v)) -> %v", rpm, ToRotationPerMinute(SI(rpm)))
	}

	if ToRotationPerMinute(rpm) != rpm {
		t.Errorf("ToRotationPerMinute(%v) -> %v", rpm, ToRotationPerMinute(rpm))
	}

	nm := NewtonMeter(107)
	if nm.SI() != 107 {
		t.Errorf("NewtonMeter(%v).SI() -> %f != 107", nm, nm.SI())
	}

	if ToNewtonMeter(SI(nm)) != nm {
		t.Errorf("ToNewtonMeter(SI(%v)) -> %v", nm, ToNewtonMeter(SI(nm)))
	}

	if ToNewtonMeter(nm) != nm {
		t.Errorf("ToNewtonMeter(%v) -> %v", nm, ToNewtonMeter(nm))
	}

	inch := Inch(10)
	if inch.SI() != 0.254 {
		t.Errorf("Inch(%v).SI() -> %f != 0.254", inch, inch.SI())
	}

	if ToInch(SI(inch)) != inch {
		t.Errorf("ToInch(SI(%v)) -> %v", inch, ToInch(SI(inch)))
	}

	if ToInch(inch) != inch {
		t.Errorf("ToInch(%v) -> %v", inch, ToInch(inch))
	}

	mm := Millimeter(3000)
	if mm.SI() != 3 {
		t.Errorf("Millimeter(%v).SI() -> %f != 3", mm, mm.SI())
	}

	if ToMillimeter(SI(mm)) != mm {
		t.Errorf("ToMillimeter(SI(%v)) -> %v", mm, ToMillimeter(SI(mm)))
	}

	if ToMillimeter(mm) != mm {
		t.Errorf("ToMillimeter(%v) -> %v", mm, ToMillimeter(mm))
	}
}
