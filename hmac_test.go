package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

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

func TestHmacSha256(t *testing.T) {
	expected, _ := hex.DecodeString("f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8")
	actual := HmacSha256([]byte("key"), []byte("The quick brown fox jumps over the lazy dog"))
	fmt.Println(expected)
	fmt.Println(actual)
	if !equal(expected, actual[:]) {
		t.Errorf("HMAC wrong!\n")
	}
}
