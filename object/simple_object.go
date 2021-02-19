package object

import (
	"github.com/benbearchen/fansica/formula"

	"fmt"
	"reflect"
)

type Namer interface {
	Name() string
	TypeName(obj Object) string
	NamedString(obj Object) string
}

type SimpleNamer struct {
	name string
}

func (namer *SimpleNamer) Name() string {
	return namer.name
}

func (namer *SimpleNamer) TypeName(obj Object) string {
	t := reflect.TypeOf(obj).Elem()
	return t.Name()
}

func (namer *SimpleNamer) NamedString(obj Object) string {
	return namer.name + "(" + namer.TypeName(obj) + ")"
}

type SimpleRotatorSocket struct {
	source Rotator
	target RotatorSocket

	rpm    formula.RotationPerMinute
	torque formula.NewtonMeter
}

func NewSimpleRotatorSocket(source Rotator) *SimpleRotatorSocket {
	return &SimpleRotatorSocket{source, nil, 0, 0}
}

func (s *SimpleRotatorSocket) String() string {
	target := ""
	if s.target != nil {
		target = "<->" + s.target.Source().String()
	}

	return fmt.Sprintf("(%s%s, %v*%v)", s.source.String(), target, s.rpm, s.torque)
}

func (s *SimpleRotatorSocket) Source() Object {
	return s.SourceRotator()
}

func (s *SimpleRotatorSocket) Target() Socket {
	return s.TargetRotator()
}

func (s *SimpleRotatorSocket) Connect(target Socket) error {
	if s.target != nil {
		return UsedSocketError
	}

	r, ok := target.(RotatorSocket)
	if !ok {
		return UnmatchSocketError
	}

	err := s.ConnectRotator(r)
	if err == nil {
		err = r.ConnectRotator(s)
	}

	return err
}

func (s *SimpleRotatorSocket) IsConnected() bool {
	return s.target != nil
}

func (s *SimpleRotatorSocket) Disband() {
	if t := s.target; t != nil {
		s.target = nil
		t.Disband()
	}
}

func (s *SimpleRotatorSocket) Power() formula.Kilowatt {
	return formula.CalcRotatorPower(s.rpm, s.torque)
}

func (s *SimpleRotatorSocket) SourceRotator() Rotator {
	return s.source
}

func (s *SimpleRotatorSocket) TargetRotator() RotatorSocket {
	return s.target
}

func (s *SimpleRotatorSocket) ConnectRotator(rs RotatorSocket) error {
	s.target = rs
	return nil
}

func (s *SimpleRotatorSocket) SetSpeedOfRatotion(rpm formula.RotationPerMinute) {
	s.rpm = rpm
}

func (s *SimpleRotatorSocket) SetTorque(t formula.NewtonMeter) {
	s.torque = t
}

func (s *SimpleRotatorSocket) RotateTorque() error {
	t := s.TargetRotator()
	t.SetSpeedOfRatotion(s.rpm)
	t.SetTorque(s.torque)

	return t.Source().InputSocket(t)
}

type SimpleElectricSocket struct {
	source Electric
	target ElectricSocket

	voltage float64
	current float64
}

func NewSimpleElectricSocket(source Electric) *SimpleElectricSocket {
	return &SimpleElectricSocket{source, nil, 0, 0}
}

func (s *SimpleElectricSocket) String() string {
	target := ""
	if s.target != nil {
		target = "<->" + s.target.Source().String()
	}

	return fmt.Sprintf("(%s%s, %.5gV*%.5gA)", s.source.String(), target, s.voltage, s.current)
}

func (s *SimpleElectricSocket) Source() Object {
	return s.SourceElectric()
}

func (s *SimpleElectricSocket) Target() Socket {
	return s.TargetElectric()
}

func (s *SimpleElectricSocket) Connect(target Socket) error {
	if s.target != nil {
		return UsedSocketError
	}

	e, ok := target.(ElectricSocket)
	if !ok {
		return UnmatchSocketError
	}

	err := s.ConnectElectric(e)
	if err == nil {
		err = e.ConnectElectric(s)
	}

	return err
}

func (s *SimpleElectricSocket) IsConnected() bool {
	return s.target != nil
}

func (s *SimpleElectricSocket) Disband() {
	if t := s.target; t != nil {
		s.target = nil
		t.Disband()
	}
}

func (s *SimpleElectricSocket) Power() formula.Kilowatt {
	w := s.voltage * s.current
	return formula.ToKilowatt(w)
}

func (s *SimpleElectricSocket) SourceElectric() Electric {
	return s.source
}

func (s *SimpleElectricSocket) TargetElectric() ElectricSocket {
	return s.target
}

func (s *SimpleElectricSocket) ConnectElectric(es ElectricSocket) error {
	s.target = es
	return nil
}

func (s *SimpleElectricSocket) SetVoltage(v float64) {
	s.voltage = v
}

func (s *SimpleElectricSocket) SetCurrent(a float64) {
	s.current = a
}
