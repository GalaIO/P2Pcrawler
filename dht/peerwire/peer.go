package peerwire

import "crypto/sha1"

func generatePeerId(str string) []byte {
	sha160 := sha1.Sum([]byte(str))
	return sha160[:]
}
