package main

import "crypto/sha256"

// HMAC This function assumed the input hash is in a clean state.
func HMAC(key, message []byte) (mac []byte) {
	// Normalize key length
	if len(key) > sha256.BlockSize {
		tmp := sha256.Sum256(key)
		key = tmp[:]
	}
	if len(key) < sha256.BlockSize {
		paddingLen := sha256.BlockSize - len(key)
		padding := make([]byte, paddingLen)
		key = append(key, padding...)
	}

	// Generate ipad and opad
	ipad := make([]byte, len(key))
	for i := range key {
		ipad[i] = key[i] ^ 0x36
	}
	opad := make([]byte, len(key))
	for i := range key {
		opad[i] = key[i] ^ 0x5c
	}

	// Perform hashing
	ihash := sha256.Sum256(append(ipad, message...))
	ohash := sha256.Sum256(append(opad, ihash[:]...))

	return ohash[:]
}
