package main

import "crypto/sha256"

func HmacSha256NormalizeKey(key []byte) []byte {
	if len(key) > sha256.BlockSize {
		tmp := sha256.Sum256(key)
		key = tmp[:]
	}
	if len(key) < sha256.BlockSize {
		paddingLen := sha256.BlockSize - len(key)
		padding := make([]byte, paddingLen)
		key = append(key, padding...)
	}
	return key
}

func HmacSha256(key, message []byte) (mac [32]byte) {
	// Normalize key length if necessary
	if len(key) != sha256.BlockSize {
		key = HmacSha256NormalizeKey(key)
	}

	// Generate ipad and opad
	ipad := make([]byte, len(key))
	opad := make([]byte, len(key))
	for i := range key {
		ipad[i] = key[i] ^ 0x36
		opad[i] = key[i] ^ 0x5c
	}

	// Perform hashing
	ihash := sha256.Sum256(append(ipad, message...))
	ohash := sha256.Sum256(append(opad, ihash[:]...))

	return ohash
}
