package object

import (
	"github.com/benbearchen/fansica/formula"

	"fmt"
)

type Motor struct {
	SimpleNamer

	maxPower  formula.Kilowatt
	maxTorque formula.NewtonMeter
	maxRPM    formula.RotationPerMinute
	eff       formula.Map

	controller bool

	socketR *SimpleRotatorSocket
	socketE *SimpleElectricSocket
}

func NewMotor(name string, power formula.Kilowatt, torque formula.NewtonMeter, rpm formula.RotationPerMinute, efficiency formula.Map) *Motor {
	// TODO:
	motor := new(Motor)
	motor.name = name
	motor.maxPower = power
	motor.maxTorque = torque
	motor.maxRPM = rpm
	motor.eff = efficiency

	motor.socketR = NewSimpleRotatorSocket(motor)
	motor.socketE = NewSimpleElectricSocket(motor)

	return motor
}

func EasyMakeMotor(name string, power, torque, rpm float64) *Motor {
	return NewMotor(name, formula.Kilowatt(power), formula.NewtonMeter(torque), formula.RotationPerMinute(rpm), nil)
}

func (motor *Motor) String() string {
	return motor.NamedString(motor)
}

func (motor *Motor) Sockets() []Socket {
	return []Socket{motor.socketR, motor.socketE}
}

func (motor *Motor) Disband() {
	motor.socketR.Disband()
	motor.socketE.Disband()
}

func (motor *Motor) InputSocket(s Socket) error {
	if s == motor.socketR {
		powerInput := motor.socketR.Power()
		if motor.socketR.rpm > motor.maxRPM || powerInput > motor.maxPower {
			return OverflowError
		}

		eff, err := motor.efficiency(motor.socketR.rpm, motor.socketR.torque)
		if err != nil {
			return err
		}

		return motor.socketE.SetPower(powerInput * formula.Kilowatt(eff))
	} else if s == motor.socketE {
		powerInput := motor.socketE.Power()
		if powerInput > motor.maxPower {
			return OverflowError
		}

		rpm := motor.socketR.rpm
		if rpm == 0 {
			return fmt.Errorf("%s.socketR.rpm == 0", motor)
		}

		t := formula.CalcRotatorTorque(powerInput, rpm)
		eff, err := motor.efficiency(rpm, t)
		if err != nil {
			return err
		}

		return motor.socketR.SetPower(powerInput * formula.Kilowatt(eff))
	} else {
		return UnmatchSocketError
	}

	return nil
}

func (motor *Motor) SetController(c bool) {
	motor.controller = c
}

func (motor *Motor) IsController() bool {
	return motor.controller
}

func (motor *Motor) efficiency(rpm formula.RotationPerMinute, t formula.NewtonMeter) (float64, error) {
	if motor.eff == nil {
		return 0.9, nil
	} else {
		return motor.eff.Value(float64(rpm), float64(t))
	}
}
