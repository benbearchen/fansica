package object

import (
	"github.com/benbearchen/fansica/formula"

	"errors"
)

type Object interface {
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

type Electric interface {
	Object

	SetVoltage(v float64)
	SetCurrent(a float64)
}

type Socket interface {
	Source() Object
	Target() Object

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
	TargetRotator() Rotator

	ConnectRotator(rs RotatorSocket) error
}

type ElectricSocket interface {
	Socket

	SourceElectric() Electric
	TargetElectric() Electric

	ConnectElectric(es ElectricSocket) error
}

func SelectRotatorSocket(a Object, index int) (RotatorSocket, error) {
	ss := a.Sockets()
	c := 0
	var first RotatorSocket
	for _, s := range ss {
		r, ok := s.(RotatorSocket)
		if !ok {
			continue
		}

		if index < 0 && first == nil {
			first = r
		}

		if c == index {
			return r, nil
		} else {
			c++
		}
	}

	if c == 0 {
		return nil, NoSocketError
	}

	if index < 0 {
		if c == 1 {
			return first, nil
		} else {
			return nil, MultiSocketError
		}
	} else {
		return nil, NoSocketError
	}
}

func ConnectRotator(a, b Object) error {
	ra, err := SelectRotatorSocket(a, -1)
	if err != nil {
		return err
	}

	rb, err := SelectRotatorSocket(b, -1)
	if err != nil {
		return err
	}

	return ra.Connect(rb)
}

func SelectElectricSocket(a Object, index int) (ElectricSocket, error) {
	ss := a.Sockets()
	c := 0
	var first ElectricSocket
	for _, s := range ss {
		e, ok := s.(ElectricSocket)
		if !ok {
			continue
		}

		if index < 0 && first == nil {
			first = e
		}

		if c == index {
			return e, nil
		} else {
			c++
		}
	}

	if c == 0 {
		return nil, NoSocketError
	}

	if index < 0 {
		if c == 1 {
			return first, nil
		} else {
			return nil, MultiSocketError
		}
	} else {
		return nil, NoSocketError
	}
}

func ConnectElectric(a, b Object, indexOfA int) error {
	ea, err := SelectElectricSocket(a, indexOfA)
	if err != nil {
		return err
	}

	eb, err := SelectElectricSocket(b, -1)
	if err != nil {
		return err
	}

	return ea.Connect(eb)
}
