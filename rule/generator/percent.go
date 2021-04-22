package generator

import "math/rand"

// p: [0, 100)
func Percent(p float64) bool {
	return rand.Float64()*100.0 <= p
}
