package main


func investmentCalc(principal float64, rate float64, time float64) float64 {
	return principal * (1 + rate*time)
}