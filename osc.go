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
	tau := math.Pi * 2
	angle := tau / sampleRate
	amplitude := 0.5
	for t := 0; t < sampleRate; t++ {
		signFunc := -1 * amplitude
		signFuncInput := math.Sin(angle * freq * float64(t))
		if signFuncInput == 0 {
			signFunc = 0
		} else if signFuncInput > 0 {
			signFunc = 1 * amplitude
		}

		samples = append(samples, float64(signFunc))
	}
	return samples
}

func sumTerms(terms []float64) float64 {
	result := float64(0)
	for _, v := range terms {
		result += v
	}
	return result
}

func FouriersBlocky(freq float64, terms int) (samples []float64) {
	tau := math.Pi * 2
	angle := tau / sampleRate
	amplitude := 0.5
	for t := 0; t < sampleRate; t++ {
		termSamples := []float64{}
		for k := 0; k < terms; k++ {
			signFunc := -1 * amplitude
			termBit := float64(2*k - 1)
			signFuncInput := math.Sin((angle * termBit * freq * float64(t)) / termBit)
			if signFuncInput > 0 {
				signFunc = 1 * amplitude
			}

			termSamples = append(termSamples, float64(signFunc))
		}
		samples = append(samples, sumTerms(termSamples))
	}
	return samples
}

func FouriersBlockyMoreRight(freq float64, terms int) (samples []float64) {
	tau := math.Pi * 2
	angle := tau / sampleRate
	amplitude := 0.5
	for t := 0; t < sampleRate; t++ {
		termSamples := []float64{}
		for k := 0; k < terms; k++ {
			signFunc := -1 * amplitude
			termBit := float64(2*k - 1)
			signFuncInput := math.Sin((angle * termBit * freq * float64(t))) / termBit
			if signFuncInput > 0 {
				signFunc = 1 * amplitude
			}

			termSamples = append(termSamples, float64(signFunc))
		}
		samples = append(samples, sumTerms(termSamples))
	}
	return samples
}

func Fouriers(freq float64, terms int) (samples []float64) {
	tau := math.Pi * 2
	angle := tau / sampleRate
	for t := 0; t < sampleRate; t++ {
		termSamples := []float64{}
		for k := 0; k < terms; k++ {
			// signFunc := -1 * amplitude
			termBit := float64(2*k - 1)
			termSample := (4 / math.Pi) * (math.Sin(angle*termBit*freq*float64(t)) / termBit)

			termSamples = append(termSamples, float64(termSample))
		}
		samples = append(samples, sumTerms(termSamples))
	}
	return samples
}

func SawTooth(freq float64, amplitude float64, phaseShift float64, verticalOffset float64) (samples []float64) {
	// angle := tau / sampleRate         // what's up with this exactly?
	period := freq
	firstBit := 4 * amplitude / period
	secondBit := period / phaseShift
	thirdBit := period / 2
	offsetBit := amplitude + verticalOffset
	for x := 0; x < sampleRate; x++ { // this looping up to sampleRate means each return is 1 second long...?
		samp := (firstBit * math.Floor(math.Mod((float64(x)-secondBit), period)-thirdBit)) - offsetBit
		samples = append(samples, float64(samp))
	}
	return samples
}

func Triangle(freq float64, amplitude float64) (samples []float64) {
	// angle := tau / sampleRate         // what's up with this exactly?
	period := freq
	firstBit := (2 * amplitude) / math.Pi
	secondBit := (2 * math.Pi) / period
	for x := 0; x < sampleRate; x++ { // this looping up to sampleRate means each return is 1 second long...?
		samp := firstBit * math.Acos(math.Sin(secondBit*float64(x)))
		samples = append(samples, float64(samp))
	}
	return samples
}
