package ForGeneration

import "crypto/sha256"

func HashSHA256(data []byte) []byte {
	h := sha256.New()
	_, _ = h.Write(data) //nolint:errcheck // No need to check error.

	return h.Sum(nil)
}
