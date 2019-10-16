package main

import "crypto/sha256"

// HMAC This function assumed the input hash is in a clean state.
func HMAC(key, message []byte) (mac []byte) {
	if len(key) > sha256.BlockSize {
		key = sha256.Sum256(key)[:]
	}

	if len(key) < sha256.BlockSize {
		paddingLen := sha256.BlockSize - len(key)
		padding := make([]byte, paddingLen)
		key = append(key, padding...)
	}

	ikp := make([]byte, len(key))
	for i := range key {
		ikp[i] = key[i] ^ 0x36
	}
	message = sha256.Sum256(append(ikp, message...))[:]

	okp := make([]byte, len(key))
	for i := range key {
		okp[i] = key[i] ^ 0x5c
	}
	mac = sha256.Sum256(append(okp, message...))[:]

	return mac
}
