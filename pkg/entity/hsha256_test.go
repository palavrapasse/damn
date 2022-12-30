package entity

import (
	"testing"
)

func TestHSHA256ComputesTheCorrectSHA256Hash(t *testing.T) {
	plainText := "import"
	plainTextSHA256Hash := "d942f64886578d8747312e368ed92d9f6b2a8d45556f0f924e2444fe911d15af"

	plainTextComputedSHA256Hash := NewHSHA256(plainText)

	if plainTextSHA256Hash != string(plainTextComputedSHA256Hash) {
		t.Fatalf(
			"Both computed hashes should be the same, however they mismatch\nPreviously Computed Hash (SHA-256): %s\nComputed Hash (SHA-256): %s\n",
			plainTextSHA256Hash,
			plainTextComputedSHA256Hash,
		)
	}
}
