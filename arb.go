// Package arb implements formulas that are used in counting arbitrage positions
package arb

// ProbFromCoefs returns probability from the sum of coefficients
func ProbFromCoefs(coefs ...float64) float64 {
	var sum float64 = 0
	for _, coef := range coefs {
		sum += 1 / coef
	}
	return sum
}

// IsArb counts coefficients and return are they in the arbitrage position or not
func IsArb(coefs ...float64) bool {
	return ProbFromCoefs(coefs...) < 1
}

// BetOnCoef returns the amount of money from the bank that is recommended to bet
// at a given coefficient in arbitrage position
func BetOnCoef(bank, coef, arb float64) float64 {
	return bank / coef / arb
}

// IncomeFromBet returns real income from bet on the chosen coefficient
func IncomeFromBet(bank, bet, coef float64) float64 {
	return bet*coef - bank
}
