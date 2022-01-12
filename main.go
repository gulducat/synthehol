package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

// todo: envelope; ADSR (attack, decay, sustain, release)
// realtime output, input
// map music notes to frequencies -- https://microtonal-guitar.com/tutorial/the-harmonic-series-musical-ratios-intervals/
// sequencer?
// filters?
// channels? mixer sum()?

const (
	sampleRate = 44100 // samples to generate, per second

	repeat = 1
)

func main() {
	// generate()

	var freq float64 = 440.0

	var f func(float64) []float64 = Sin
	var samples []float64

	root := f(freq)
	third := f(freq * 5 / 4)
	fifth := f(freq * 3 / 2)
	octave := f(freq * 2)
	bass := f(freq / 2)
	bass2 := f(freq / 2 / 2)
	seventh := multiply(f(freq*15/8), 0.5)

	samples = root
	// chords :)
	samples = append(samples, sum(root, third)...)
	samples = append(samples, sum(root, third, fifth)...)
	samples = append(samples, sum(root, third, fifth, octave)...)
	samples = append(samples, sum(root, third, fifth, octave, bass)...)
	samples = append(samples, sum(root, third, fifth, octave, bass2)...)
	samples = append(samples, sum(root, third, fifth, octave, bass2, seventh)...)
	samples = append(samples, sum(root, third, fifth, octave, bass2, seventh)...)

	// samples = sum(
	// 	Sin(freq/2/2), // bass
	// 	Sin(freq/2),
	// 	Sin(freq),
	// 	// Sin(freq*5/4),  // major 3rd
	// 	// Sin(freq*3/2), // perfect 5th
	// 	multiply(
	// 		Sin(freq*15/8), // major 7th lel
	// 		0.25,           // is vv dissonant, so be quieter
	// 	),
	// 	// Sin(freq*2), // octave
	// )
	// samples = multiply(samples, 1000) // squareize everything.

	// var (
	// 	start float64 = 1.0
	// 	// end   float64 = 1.0e-4
	// )
	// decayfac := math.Pow(end/start, 1.0/float64(sampleRate))
	// decayfac := 0.98
	// var decayfac float64 = sampleRate / 2

	// cut all above 1 and below -1
	samples = limit(samples)

	var total []float64
	for x := 0; x < repeat; x++ {
		for _, samp := range samples {
			// samp *= start
			// start *= decayfac
			fmt.Printf("%.8f\n", samp)

			total = append(total, samp)
		}
	}
	err := WriteBin("out.bin", total)
	if err != nil {
		panic(err)
	}
}

func WriteBin(filename string, samples []float64) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	for _, sample := range samples {
		var buf [8]byte
		binary.LittleEndian.PutUint32(
			buf[:],
			math.Float32bits(float32(sample)), // todo: we're 32 bits now?
		)
		_, err := f.Write(buf[:])
		if err != nil {
			return err
		}
	}
	return nil
}

func sum(sampleSets ...[]float64) (summed []float64) {
	numSets := float64(len(sampleSets))
	numSamples := len(sampleSets[0])
	for x := 0; x < numSamples; x++ {
		summed = append(summed, 0)
		for _, set := range sampleSets {
			summed[x] += set[x] / numSets // divide by numSets to equalize all the samples...
		}
	}
	return summed
}

func sum2(a, b []float64) (sums []float64) {
	for i, n := range a {
		sums = append(sums, n+b[i])
	}
	return sums
}
