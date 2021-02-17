package object

import (
	"github.com/benbearchen/fansica/formula"
)

type Engine struct {
	SimpleNamer
	SimpleRotator

	minRPM    formula.RotationPerMinute
	maxRPM    formula.RotationPerMinute
	maxTorque formula.NewtonMeter
	maxPower  formula.Kilowatt
	eff       formula.Map

	controller bool

	socket *SimpleRotatorSocket
}

func NewEngine(name string, minRPM, maxRPM formula.RotationPerMinute, torque formula.NewtonMeter, power formula.Kilowatt, eff formula.Map) *Engine {
	eng := new(Engine)
	eng.name = name
	eng.minRPM = minRPM
	eng.maxRPM = maxRPM
	eng.maxTorque = torque
	eng.maxPower = power
	eng.eff = eff

	eng.socket = NewSimpleRotatorSocket(eng)

	return eng
}

func EasyMakeEngine(name string, maxRPM, torque, power float64) *Engine {
	return NewEngine(name, formula.RotationPerMinute(800), formula.RotationPerMinute(maxRPM), formula.NewtonMeter(torque), formula.Kilowatt(power), nil)
}

func (eng *Engine) String() string {
	return eng.NamedString(eng)
}

func (eng *Engine) Sockets() []Socket {
	return []Socket{eng.socket}
}

func (eng *Engine) Disband() {
	eng.socket.Disband()
}

func (eng *Engine) SetController(c bool) {
	eng.controller = c
}

func (eng *Engine) IsController() bool {
	return eng.controller
}
