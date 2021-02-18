package object

import "testing"

func TestCar(t *testing.T) {
	b := EasyMakeBattery("b0", 40, 220)
	eng := EasyMakeEngine("ice", 6500, 220, 96)
	mg0 := EasyMakeMotor("m0", 100, 240, 12000)
	sr0 := NewSingleReducer("sr0", 1/1.8)
	mg1 := EasyMakeMotor("m1", 100, 240, 12000)
	sr1 := NewSingleReducer("sr1", 10)
	mg2 := EasyMakeMotor("m2", 140, 290, 12000)
	sr2 := NewSingleReducer("sr2", 10)
	wheel1, _ := ParseWheel("wf", "255/50R20")
	wheel2, _ := ParseWheel("wr", "255/50R20")
	pcu := NewPCU("pcu", 4)

	chanRotator := func(a, b Rotator, c Reducer) {
		err := ChanRotator(a, b, c)
		if err != nil {
			t.Errorf("ChanRotator(%v, %v, %v) failed: %v", a, b, c, err)
		}
	}

	chanRotator(eng, mg0, sr0)
	chanRotator(mg1, wheel1, sr1)
	chanRotator(mg2, wheel2, sr2)

	connectElectric := func(a, b Electric, i int) {
		err := ConnectElectric(a, b, i)
		if err != nil {
			t.Errorf("ConnectElectric(%v, %v) failed: %v", a, b, err)
		}
	}

	connectElectric(pcu, mg0, 0)
	connectElectric(pcu, mg1, 1)
	connectElectric(pcu, mg2, 2)
	connectElectric(pcu, b, 3)

	checkRotatorChan := func(a Object, c int, b Object) {
		cnt, obj := CountRotatorChanController(a)
		if cnt != c || obj != b {
			t.Errorf("CountRotatorChanController(%v) -> (%d, %v) != (%d, %v)", a, cnt, obj, c, b)
		}
	}

	checkRotatorChan(eng, 0, nil)
	checkRotatorChan(sr0, 0, nil)
	checkRotatorChan(mg0, 0, nil)
	checkRotatorChan(mg1, 0, nil)
	checkRotatorChan(sr1, 0, nil)
	checkRotatorChan(wheel1, 0, nil)
	checkRotatorChan(mg2, 0, nil)
	checkRotatorChan(sr2, 0, nil)
	checkRotatorChan(wheel2, 0, nil)

	eng.SetController(true)
	wheel1.SetController(true)
	wheel2.SetController(true)

	checkRotatorChan(eng, 1, eng)
	checkRotatorChan(sr0, 1, eng)
	checkRotatorChan(mg0, 1, eng)
	checkRotatorChan(mg1, 1, wheel1)
	checkRotatorChan(sr1, 1, wheel1)
	checkRotatorChan(wheel1, 1, wheel1)
	checkRotatorChan(mg2, 1, wheel2)
	checkRotatorChan(sr2, 1, wheel2)
	checkRotatorChan(wheel2, 1, wheel2)
}
