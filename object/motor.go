package object

import (
	"github.com/benbearchen/fansica/formula"
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
		if motor.socketR.rpm > motor.maxRPM || motor.socketR.Power() > motor.maxPower {
			return OverflowError
		}

		// TODO: 输出功率？
	} else if s == motor.socketE {
		if motor.socketE.Power() > motor.maxPower {
			return OverflowError
		}

		// TODO: 输出转速？？
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
