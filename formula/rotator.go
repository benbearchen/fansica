package formula

import (
	"math"
)

const (
	RotatorPowerRatio = 60000 / 2 / math.Pi
)

func CalcRotatorPower(rpm RotationPerMinute, t NewtonMeter) Kilowatt {
	return Kilowatt(float64(rpm) * float64(t) / RotatorPowerRatio)
}

func CalcRotatorTorque(kw Kilowatt, rpm RotationPerMinute) NewtonMeter {
	return NewtonMeter(float64(kw) * RotatorPowerRatio / float64(rpm))
}
