package main

import (
	"crypto/sha256"
	"testing"

	"golang.org/x/crypto/pbkdf2"
)

func TestPbkdf2Sha256(t *testing.T) {
	password := []byte("lepassword")
	salt := []byte("lesalt")
	iterations := 50000
	keyLen := 32
	expected := pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
	actual := Pbkdf2HmacSha256(password, salt, iterations, keyLen)
	if !equal(expected, actual) {
		t.Errorf("PBKDF2 wrong!\n%v expected,\n%v actual", expected, actual)
	}
}
