package object

import (
	"github.com/benbearchen/fansica/formula"

	"fmt"
	"math"
)

type Wheel struct {
	SimpleRotator

	r       formula.Inch
	width   formula.Millimeter
	profile float64 // 扁平率

	controller bool

	socket *SimpleRotatorSocket
}

func ParseWheel(spec string) (*Wheel, error) {
	var w, p, r int
	_, err := fmt.Sscanf(spec, "%d/%dR%d", &w, &p, &r)
	if err != nil {
		return nil, err
	}

	wheel := new(Wheel)
	wheel.r = formula.Inch(r)
	wheel.width = formula.Millimeter(w)
	wheel.profile = float64(p)

	wheel.socket = NewSimpleRotatorSocket(wheel)

	return wheel, nil
}

func (wheel *Wheel) Diameter() float64 {
	return formula.SI(wheel.r) + formula.SI(wheel.width)*wheel.profile/100*2
}

func (wheel *Wheel) Perimeter() float64 {
	return wheel.Diameter() * math.Pi
}

func (wheel *Wheel) Sockets() []Socket {
	return []Socket{wheel.socket}
}

func (wheel *Wheel) Disband() {
	wheel.socket.Disband()
}

func (wheel *Wheel) SetController(c bool) {
	wheel.controller = c
}

func (wheel *Wheel) IsController() bool {
	return wheel.controller
}

func (wheel *Wheel) Speed() formula.KilometerPerHour {
	s := wheel.Perimeter() * formula.SI(wheel.rpm)
	return formula.ToKilometerPerHour(s)
}
