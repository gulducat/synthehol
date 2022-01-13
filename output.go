package main

import (
	"encoding/binary"
	"math"
	"os"
)

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
