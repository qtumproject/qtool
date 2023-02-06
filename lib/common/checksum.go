package common

import (
	"crypto/sha256"
)

// Checksum() generates a 4 bytes double sha256 checksum from given input
func Checksum(input []byte) (cksum [4]byte) {
	firstSHA := sha256.Sum256(input)
	secondSHA := sha256.Sum256(firstSHA[:])
	copy(cksum[:], secondSHA[:4])
	return
}
