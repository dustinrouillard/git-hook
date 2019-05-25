package github

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

// VerifySignature will compare the incoming signature
func VerifySignature(Secret string, Payload []byte, Signature string) bool {
	if len(Signature) == 0 {
		return false
	}

	DecodedSignature, DecodeErr := hex.DecodeString(Signature[5:])
	if DecodeErr != nil {
		return false
	}

	mac := hmac.New(sha1.New, []byte(Secret))
	mac.Write(Payload)
	expected := mac.Sum(nil)

	return hmac.Equal(DecodedSignature, expected)
}

// Verify will take in the incoming github hook and verify the secret and details
func Verify(Secret string, Signature string, Payload []byte) error {
	if len(Signature) == 0 {
		return errors.New("no_signature")
	}

	// Compare signature
	if !VerifySignature(Secret, Payload, Signature) {
		return errors.New("invalid_signature")
	}

	return nil
}
