package object

import (
	"github.com/benbearchen/fansica/formula"
)

type Motor struct {
	SimpleRotator
	SimpleElectric

	maxPower  formula.Kilowatt
	maxTorque formula.NewtonMeter
	maxRPM    formula.RotationPerMinute
	eff       formula.Map

	controller bool

	socketR *SimpleRotatorSocket
	socketE *SimpleElectricSocket
}

func NewMotor(power formula.Kilowatt, torque formula.NewtonMeter, rpm formula.RotationPerMinute, efficiency formula.Map) *Motor {
	motor := new(Motor)
	motor.maxPower = power
	motor.maxTorque = torque
	motor.maxRPM = rpm
	motor.eff = efficiency

	motor.socketR = NewSimpleRotatorSocket(motor)
	motor.socketE = NewSimpleElectricSocket(motor)

	return motor
}

func EasyMakeMotor(power, torque, rpm float64) *Motor {
	return NewMotor(formula.Kilowatt(power), formula.NewtonMeter(torque), formula.RotationPerMinute(rpm), nil)
}

func (motor *Motor) Sockets() []Socket {
	return []Socket{motor.socketR, motor.socketE}
}

func (motor *Motor) Disband() {
	motor.socketR.Disband()
	motor.socketE.Disband()
}

func (motor *Motor) SetController(c bool) {
	motor.controller = c
}

func (motor *Motor) IsController() bool {
	return motor.controller
}
