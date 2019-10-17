package main

import (
	"encoding/binary"
	"math"
)

func Pbkdf2HmacSha256(password, salt []byte, iterations, outputLen int) (derivedKey []byte) {
	// Calculate how many blocks are necessary
	blocks := int(math.Ceil(float64(outputLen) / 32))

	// Allocate output array
	derivedKey = make([]byte, blocks*32)
	for b := 0; b < blocks; b++ {
		// Do all the hashing
		var bBytes [4]byte
		binary.BigEndian.PutUint32(bBytes[:], uint32(b+1))
		state := make([][]byte, iterations)
		state[0] = HmacSha256(password, append(salt, bBytes[:]...))
		for i := 1; i < iterations; i++ {
			state[i] = HmacSha256(password, state[i-1])
		}
		// XOR all the results together
		copy(derivedKey[32*b:32*b+32], state[0][:])
		for i := 1; i < len(state); i++ {
			for bi := range state[i] {
				derivedKey[32*b+bi] ^= state[i][bi]
			}
		}
	}
	return derivedKey[:outputLen]
}
