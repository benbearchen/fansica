package formula

func CalcThermalEfficiencyAdvance(cLow, wLow, cHi, wHi, charge float64) float64 {
	total := wLow / cLow * cHi
	t := total / wHi
	direct := t * wLow
	save := (total - t*wLow) * charge
	valid := direct + save
	return valid / wLow
}
