package main

import "math"

func Sin(freq float64) (samples []float64) {
	tau := math.Pi * 2
	angle := tau / sampleRate         // what's up with this exactly?
	for i := 0; i < sampleRate; i++ { // this looping up to sampleRate means each return is 1 second long...?
		samp := math.Sin(angle * freq * float64(i))
		samples = append(samples, samp)
	}
	return samples
}

func DumbSquare(freq float64) (samples []float64) {
	// start with a sine wave
	samples = Sin(freq)
	// amp it way up
	samples = multiply(samples, 1000)
	// crush back down to between -1 and 1
	return limit(samples)
	// return limit(multiply(Sin(freq), 1000))
}

// todo: fix this, it doesn't work right.
func Square(freq float64) (samples []float64) {
	half := sampleRate / 2
	for i := 0; i < sampleRate; i++ {
		if i < half {
			// if sampleRate%i < freq {
			samples = append(samples, 1)
		} else {
			samples = append(samples, -1)
		}
	}
	return samples
}
