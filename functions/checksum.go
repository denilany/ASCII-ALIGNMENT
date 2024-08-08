package functions

import (
	"crypto/sha256"
	"encoding/hex"
)

// Function checkBanner implements sha256 file encryption system to
// ensure correctness, authenticity and intergrity of every banner file.
func checkBanner(byteFile []byte) string {
	hasher := sha256.New()
	hasher.Write(byteFile)
	hashInBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}
