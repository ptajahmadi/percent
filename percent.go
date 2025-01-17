package percent

// Percent - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
// ex. 20% of 200 is 40
func Percent(percent int, all int) float64 {
	return ((float64(all) * float64(percent)) / float64(100))
}

// PercentFloat - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func PercentFloat(percent float64, all float64) float64 {
	return ((float64(all) * float64(percent)) / float64(100))
}

// PercentOf - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func PercentOf(part int, total int) float64 {
	return (float64(part) * float64(100)) / float64(total)
}

// PercentOfFloat - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func PercentOfFloat(part float64, total float64) float64 {
	return (float64(part) * float64(100)) / float64(total)
}

// Change - calculate the percent increase/decrease from two numbers.
// ex. 60 is a 200.0% increase from 20
func Change(before int, after int) float64 {
	return 100 * ((float64(after) - float64(before)) / float64(before))
}

// ChangeFloat - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0
func ChangeFloat(before float64, after float64) float64 {
	return 100 * ((after - before) / float64(before))
}
