package main

import (
	"math"
)

func Pbkdf2HmacSha256(password, salt []byte, iterations, outputLen int) (derivedKey []byte) {
	// Calculate how many blocks are necessary
	blocks := int(math.Ceil(float64(outputLen) / 32))

	// Allocate variables
	derivedKey = make([]byte, blocks*32)
	state := make([][]byte, iterations)

	for b := 0; b < blocks; b++ {
		// Do all the hashing
		state[0] = HmacSha256(password, append(salt, uint32Bytes(b+1)...))
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
