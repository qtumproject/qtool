package common

import (
	"crypto/sha256"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

// HashPubKey() hashes twice a public key and returns a RIPEMD160 hash
//
// If the recevied public key is not 33 bytes, an error is returned
func HashPubKey(pubKey []byte) ([]byte, error) {
	if len(pubKey) != 33 {
		return nil, errors.Errorf("pubKey must be 33 bytes, but is %d", len(pubKey))
	}
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		return nil, err
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160, nil
}
