package object

import (
	"github.com/benbearchen/fansica/formula"

	"errors"
)

type Object interface {
	Name() string
	String() string

	Sockets() []Socket
	Disband()

	SetController(c bool)
	IsController() bool
}

type Rotator interface {
	Object

	SetSpeedOfRatotion(rpm formula.RotationPerMinute)
	SetTorque(t formula.NewtonMeter)
}

type Reducer interface {
	Object

	Ratio() float64
}

type Transmission interface {
	Reducer

	RatioRange() (float64, float64)
	SetRatio(r float64)
}

type Clutch interface {
	Engage(engage bool)
}

type Electric interface {
	Object

	SetVoltage(v float64)
	SetCurrent(a float64)
}

type Socket interface {
	Source() Object
	Target() Socket

	Connect(s Socket) error
	IsConnected() bool
	Disband()
}

var (
	NoSocketError      = errors.New("no socket")
	UnmatchSocketError = errors.New("unmatch socket")
	UsedSocketError    = errors.New("used socket")
	MultiSocketError   = errors.New("multi socket")
)

type RotatorSocket interface {
	Socket

	SourceRotator() Rotator
	TargetRotator() RotatorSocket

	ConnectRotator(rs RotatorSocket) error
}

type ElectricSocket interface {
	Socket

	SourceElectric() Electric
	TargetElectric() ElectricSocket

	ConnectElectric(es ElectricSocket) error
}
