package object

import "testing"

func TestCar(t *testing.T) {
	b := EasyMakeBattery(40, 220)
	eng := EasyMakeEngine(6500, 220, 96)
	mg0 := EasyMakeMotor(100, 240, 12000)
	mg1 := EasyMakeMotor(100, 240, 12000)
	mg2 := EasyMakeMotor(140, 290, 12000)
	wheel1, _ := ParseWheel("255/50R20")
	wheel2, _ := ParseWheel("255/50R20")
	pcu := NewPCU(4)

	connectRotator := func(a, b Rotator) {
		err := ConnectRotator(a, b)
		if err != nil {
			t.Errorf("ConnectRotator(%v, %v) failed: %v", a, b, err)
		}
	}

	connectRotator(eng, mg0)
	connectRotator(mg1, wheel1)
	connectRotator(mg2, wheel2)

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
}
