package tools

import (
	"crypto/sha256"
)

// TODO: delete this file. Replaced by the one in the crypto package

// Checksum() generates a 4 bytes double sha256 checksum from given input
func Checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}
