package main

// limit clips all values above 1 or below -1
func limit(samples []float64) (clipped []float64) {
	for _, s := range samples {
		c := s
		if s > 1 {
			c = 1
		} else if s < -1 {
			c = -1
		}
		clipped = append(clipped, c)
	}
	return clipped
}

// multiply all values by factor
func multiply(samples []float64, factor float64) (out []float64) {
	for _, s := range samples {
		out = append(out, s*factor)
	}
	return out
}

// func filter(samples []float64, filter func(float64) float64) (filtered []float64) {
// 	for _, s := range samples {
// 		filtered = append(filtered, filter(s))
// 	}
// 	return filtered
// }

// func divide(x float64) func(float64) float64 {
// 	return func(val float64) float64 {
// 		return val / x
// 	}
// }

// func multiply(x float64) func(float64) float64 {
// 	return func(val float64) float64 {
// 		return val * x
// 	}
// }
