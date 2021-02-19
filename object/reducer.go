package object

import (
	"github.com/benbearchen/fansica/formula"
)

type SingleReducer struct {
	SimpleNamer

	ratio float64

	sockets [2]*SimpleRotatorSocket
}

func NewSingleReducer(name string, ratio float64) *SingleReducer {
	r := new(SingleReducer)
	r.name = name
	r.ratio = ratio
	r.sockets[0] = NewSimpleRotatorSocket(r)
	r.sockets[1] = NewSimpleRotatorSocket(r)

	return r
}

func (r *SingleReducer) String() string {
	return r.NamedString(r)
}

func (r *SingleReducer) Sockets() []Socket {
	return []Socket{r.sockets[0], r.sockets[1]}
}

func (r *SingleReducer) Disband() {
	r.sockets[0].Disband()
	r.sockets[1].Disband()
}

func (r *SingleReducer) InputSocket(s Socket) error {
	// TODO: 扭矩方向？？
	eff := 0.99
	s0, s1 := r.sockets[0], r.sockets[1]
	if s0 == s {
		s1.SetSpeedOfRatotion(formula.RotationPerMinute(float64(s0.rpm) / r.Ratio()))
		s1.SetTorque(formula.NewtonMeter(float64(s0.torque) * r.Ratio() * eff))
	} else if s1 == s {
		s0.SetSpeedOfRatotion(formula.RotationPerMinute(float64(s1.rpm) * r.Ratio()))
		s0.SetTorque(formula.NewtonMeter(float64(s1.torque) / r.Ratio() / eff))
	} else {
		return UnmatchSocketError
	}

	return nil
}

func (r *SingleReducer) SetController(c bool) {
	// 不会成为控制焦点
}

func (r *SingleReducer) IsController() bool {
	return false
}

func (r *SingleReducer) Ratio() float64 {
	return r.ratio
}
