package object

import (
	"github.com/benbearchen/fansica/formula"
)

type Engine struct {
	SimpleNamer

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

func (eng *Engine) InputSocket(s Socket) error {
	if s != eng.socket {
		return UnmatchSocketError
	}

	if eng.socket.Power() > eng.maxPower || eng.socket.rpm > eng.maxRPM || eng.socket.rpm < eng.minRPM {
		return OverflowError
	}

	return nil
}

func (eng *Engine) SetController(c bool) {
	eng.controller = c
}

func (eng *Engine) IsController() bool {
	return eng.controller
}

func (eng *Engine) SetSpeedOfRatotion(rpm formula.RotationPerMinute) {
	eng.socket.SetSpeedOfRatotion(rpm)
}

func (eng *Engine) SetTorque(torque formula.NewtonMeter) {
	eng.socket.SetTorque(torque)
}
