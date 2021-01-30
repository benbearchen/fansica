package formula

func CalcGasoline100km(power, bsfc, v float64) float64 {
	kw := float64(ToKilowatt(power))
	kph := float64(ToKilometerPerHour(v))
	fc := kw / bsfc / 9 * 100 / kph
	return SI(LiterPer100Kilometer(fc))
}
