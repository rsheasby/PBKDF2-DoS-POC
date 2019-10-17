package main

import "encoding/binary"

func uint32Bytes(i int) []byte {
	var bBytes [4]byte
	binary.BigEndian.PutUint32(bBytes[:], uint32(i))
	return bBytes[:]
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
