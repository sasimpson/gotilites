package maths

import "math"

//RoundInt - rounds integers to a place:
// Example: RoundInt(12345, 3) would return 12000
func RoundInt(num, places int) int {
	_num := float64(num)
	factor := int(math.Pow(10, float64(places)))
	_div := int(math.Abs(_num)) / factor
	_mod := int(math.Abs(_num)) % factor
	if (_mod) >= (factor / 2) {
		return int(math.Copysign(float64(((_div + 1) * factor)), _num))
	}
	return int(math.Copysign(float64(((_div) * factor)), _num))
}
