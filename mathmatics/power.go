package mathmatics

import "math"

func PowerEhsan(base, power int) int {
	sum := 1
	for i := 0; i < power; i++ {
		sum *= base
	}
	return sum

}
func PowerGo(base, power float64) float64 {

	return (math.Pow((base), (power)))

}
