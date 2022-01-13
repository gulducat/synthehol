package main

import (
	"encoding/binary"
	"math"
	"os"

	"github.com/gordonklaus/portaudio"
)

func StreamOut(samples []float64) (*portaudio.Stream, error) {
	if err := portaudio.Initialize(); err != nil {
		return nil, err
	}
	defer portaudio.Terminate()

	chunks := 8192
	buf := make([]float32, chunks)

	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, chunks, &buf)
	if err != nil {
		return stream, err
	}
	if err := stream.Start(); err != nil {
		return stream, err
	}
	defer stream.Stop()

	for i, s := range samples {
		idx := i % chunks
		// buf = append(buf, float32(s)) // iiinteresting click tempo thing goin here...
		if idx == 0 {
			if err := stream.Write(); err != nil {
				return stream, err
			}
			buf = make([]float32, chunks)
		}
		buf[idx] = float32(s)
	}

	return stream, nil
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
