package main

import "testing"

func TestCalcPie(t *testing.T) {
	for i := 1; i < 50000; i++ {
		result := calcPie(i, i+1)
		if result != i*(i+1) {
			t.Errorf("wrong Data")
		}
	}
}
