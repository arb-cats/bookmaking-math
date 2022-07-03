package arb

import (
	"testing"
)

func TestIsArb(t *testing.T) {
	type args struct {
		coefs []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "arb from 2 coefficients", args: args{coefs: []float64{2.05, 2.10}}, want: true},
		{name: "not an arb from 2 coefficients", args: args{coefs: []float64{1.9, 2.01}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArb(tt.args.coefs...); got != tt.want {
				t.Errorf("IsArb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProbFromCoefs(t *testing.T) {
	type args struct {
		coefs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "sum 2 coefficients higher than 2", args: args{coefs: []float64{2.05, 2.10}}, want: 0.9639953542392568},
		{name: "sum 2 coefficients lover than 2", args: args{coefs: []float64{1.9, 1.99}}, want: 1.0288283522877544},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProbFromCoefs(tt.args.coefs...); got != tt.want {
				t.Errorf("ProbFromCoefs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetOnCoef(t *testing.T) {
	type args struct {
		bank float64
		coef float64
		arb  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "count bank for correct bet", args: args{bank: 1000, coef: 2.1, arb: 0.95}, want: 501.25313283208015},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetOnCoef(tt.args.bank, tt.args.coef, tt.args.arb); got != tt.want {
				t.Errorf("BetOnCoef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncomeFromBet(t *testing.T) {
	type args struct {
		bank float64
		bet  float64
		coef float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "positive income", args: args{bank: 50000, bet: 16000, coef: 3.2}, want: 1200},
		{name: "negative income", args: args{bank: 50000, bet: 16000, coef: 2.2}, want: -14800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncomeFromBet(tt.args.bank, tt.args.bet, tt.args.coef); got != tt.want {
				t.Errorf("IncomeFromBet() = %v, want %v", got, tt.want)
			}
		})
	}
}
