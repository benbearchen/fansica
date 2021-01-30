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
}
