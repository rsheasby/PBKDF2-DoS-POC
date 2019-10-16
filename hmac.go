package main

import "hash"

// HMAC This function assumed the input hash is in a clean state.
func HMAC(key, message []byte, hash hash.Hash) (mac []byte) {
	if len(key) > hash.BlockSize() {
		hash.Write(key)
		key = hash.Sum([]byte{})
		hash.Reset()
	} else if len(key) < hash.BlockSize() {
		padding := make([]byte, hash.Size()-len(key))
		append(key, padding)
	}

	ikp := make([]byte, len(key))
	for i, v := range key {
		ikp[i] = key[i] ^ 0x36
	}
	hash.Write(append(ikp, message))
	message = hash.Sum([]byte{})
	hash.Reset()

	opk := make([]byte, len(key))
	for i, v := range key {
		okp[i] = key[i] ^ 0x5c
	}
	hash.Write(append(okp, message))
	mac = hash.Sum([]byte{})
	hash.Reset()
}
