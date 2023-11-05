package mathmatics

import (
	"testing"
)

func TestPower(t *testing.T) {
	type test struct {
		Base           int
		Power          int
		ExpectedResult int
	}
	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
	}
	//{ base: 2,Power: 0,expectedResult: 0  },
	//{ base: 2,Power: 0,expectedResult: 0  },
	for _, v := range testCases {
		res := PowerEhsan(v.Base, v.Power)
		if res != v.ExpectedResult {
			t.Errorf("not good result")

		}
	}

}

// benchmark
var result int

func BenchmarkPowerEhsan(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		res := PowerEhsan(2, 6)
		result = res

	}
}

var result2 float64

func BenchmarkPowerMath(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		res := PowerGo(2, 6)
		result2 = res

	}
}
