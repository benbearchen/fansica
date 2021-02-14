package formula

import (
	"fmt"
)

type ToSI interface {
	SI() float64
}

func SI(v interface{}) float64 {
	switch v := v.(type) {
	case ToSI:
		return v.SI() // 国际单位
	case int:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		panic(fmt.Sprintf("can't convert %v to SI float64", v))
	}
}

// 车速：千米每小时
type KilometerPerHour float64

func (kph KilometerPerHour) SI() float64 {
	return float64(kph) / 3.6
}

func (kph KilometerPerHour) String() string {
	return fmt.Sprintf("%.4gkph", float64(kph))
}

func ToKilometerPerHour(v float64) KilometerPerHour {
	return KilometerPerHour(v * 3.6)
}

// 功率：千瓦
type Kilowatt float64

func (kw Kilowatt) SI() float64 {
	return float64(kw) * 1000
}

func (kw Kilowatt) String() string {
	return fmt.Sprintf("%.5gkW", float64(kw))
}

func ToKilowatt(w interface{}) Kilowatt {
	return Kilowatt(SI(w) / 1000)
}

// 功率：英制马力
type HorsePower float64

func (hp HorsePower) SI() float64 {
	return float64(hp) * 745.7
}

func (hp HorsePower) String() string {
	return fmt.Sprintf("%.5gHP", float64(hp))
}

func ToHorsePower(w interface{}) HorsePower {
	return HorsePower(SI(w) / 745.7)
}

// 功率：公制马力
type PS float64

func (ps PS) SI() float64 {
	return float64(ps) * 735.5
}

func (ps PS) String() string {
	return fmt.Sprintf("%.5gPS", float64(ps))
}

func ToPS(w interface{}) PS {
	return PS(SI(w) / 735.5)
}

// 能量：千瓦时
type KilowattHour float64

func (kWh KilowattHour) SI() float64 {
	return float64(kWh) * 3600 * 1000
}

func (kWh KilowattHour) String() string {
	return fmt.Sprintf("%.5gkWh", float64(kWh))
}

func ToKilowattHour(j interface{}) KilowattHour {
	return KilowattHour(SI(j)) / 3600 / 1000
}

// 油耗：升/百公里
type LiterPer100Kilometer float64

func (lphkm LiterPer100Kilometer) SI() float64 {
	return float64(lphkm) / 1000 * (100 * 1000)
}

func (lphkm LiterPer100Kilometer) String() string {
	return fmt.Sprintf("%.5gL/100km", float64(lphkm))
}

func ToLiterPer100Kilometer(v interface{}) LiterPer100Kilometer {
	return LiterPer100Kilometer(SI(v)) * 1000 / (100 * 1000)
}

// 转速：转/分钟
type RotationPerMinute float64

func (rpm RotationPerMinute) SI() float64 {
	return float64(rpm) / 60
}

func (rpm RotationPerMinute) String() string {
	return fmt.Sprintf("%.5gr/min", float64(rpm))
}

func ToRotationPerMinute(v interface{}) RotationPerMinute {
	return RotationPerMinute(SI(v) * 60)
}

// 扭矩：牛米
type NewtonMeter float64

func (nm NewtonMeter) SI() float64 {
	return float64(nm)
}

func (nm NewtonMeter) String() string {
	return fmt.Sprintf("%.5gNm", float64(nm))
}

func ToNewtonMeter(v interface{}) NewtonMeter {
	return NewtonMeter(SI(v))
}

// 长度：英寸
type Inch float64

func (inch Inch) SI() float64 {
	return float64(inch) * 0.0254
}

func (inch Inch) String() string {
	return fmt.Sprintf("%.4gin", float64(inch))
}

func ToInch(v interface{}) Inch {
	return Inch(SI(v)) / 0.0254
}

// 长度：毫米
type Millimeter float64

func (mm Millimeter) SI() float64 {
	return float64(mm) / 1000
}

func (mm Millimeter) String() string {
	return fmt.Sprintf("%.5gmm", float64(mm))
}

func ToMillimeter(v interface{}) Millimeter {
	return Millimeter(SI(v) * 1000)
}
