package object

import (
	"github.com/benbearchen/fansica/formula"
)

type Battery struct {
	SimpleElectric

	capacity formula.KilowattHour
	maxPower formula.Kilowatt
	powerMap formula.Map

	socket *SimpleElectricSocket
}

func NewBattery(capacity formula.KilowattHour, power formula.Kilowatt, powerMap formula.Map) *Battery {
	b := new(Battery)
	b.capacity = capacity
	b.maxPower = power
	b.powerMap = powerMap

	b.socket = NewSimpleElectricSocket(b)

	return b
}

func EasyMakeBattery(capacity, power float64) *Battery {
	return NewBattery(formula.KilowattHour(capacity), formula.Kilowatt(power), nil)
}

func (b *Battery) Sockets() []Socket {
	return []Socket{b.socket}
}

func (b *Battery) Disband() {
	b.socket.Disband()
}

func (b *Battery) SetController(c bool) {
	// 假定电池不会成为控制焦点
}

func (b *Battery) IsController() bool {
	return false
}
