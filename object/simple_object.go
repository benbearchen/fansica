package object

import (
	"github.com/benbearchen/fansica/formula"

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

type SimpleRotator struct {
	rpm    formula.RotationPerMinute
	torque formula.NewtonMeter
}

func (r *SimpleRotator) SetSpeedOfRatotion(rpm formula.RotationPerMinute) {
	r.rpm = rpm
}

func (r *SimpleRotator) SetTorque(t formula.NewtonMeter) {
	r.torque = t
}

func (r *SimpleRotator) Power() formula.Kilowatt {
	return formula.CalcRotatorPower(r.rpm, r.torque)
}

type SimpleRotatorSocket struct {
	source Rotator
	target RotatorSocket
}

func NewSimpleRotatorSocket(source Rotator) *SimpleRotatorSocket {
	return &SimpleRotatorSocket{source, nil}
}

func (s *SimpleRotatorSocket) Source() Object {
	return s.SourceRotator()
}

func (s *SimpleRotatorSocket) Target() Object {
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

func (s *SimpleRotatorSocket) SourceRotator() Rotator {
	return s.source
}

func (s *SimpleRotatorSocket) TargetRotator() Rotator {
	return s.target.Source().(Rotator)
}

func (s *SimpleRotatorSocket) ConnectRotator(rs RotatorSocket) error {
	s.target = rs
	return nil
}

type SimpleElectric struct {
	voltage float64
	current float64
}

func (e *SimpleElectric) SetVoltage(v float64) {
	e.voltage = v
}

func (e *SimpleElectric) SetCurrent(a float64) {
	e.current = a
}

func (e *SimpleElectric) Power() formula.Kilowatt {
	w := e.voltage * e.current
	return formula.ToKilowatt(w)
}

type SimpleElectricSocket struct {
	source Electric
	target ElectricSocket
}

func NewSimpleElectricSocket(source Electric) *SimpleElectricSocket {
	return &SimpleElectricSocket{source, nil}
}

func (s *SimpleElectricSocket) Source() Object {
	return s.SourceElectric()
}

func (s *SimpleElectricSocket) Target() Object {
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

func (s *SimpleElectricSocket) SourceElectric() Electric {
	return s.source
}

func (s *SimpleElectricSocket) TargetElectric() Electric {
	return s.target.Source().(Electric)
}

func (s *SimpleElectricSocket) ConnectElectric(es ElectricSocket) error {
	s.target = es
	return nil
}
