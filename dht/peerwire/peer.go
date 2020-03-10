package peerwire

import "crypto/sha1"

func GeneratePeerId(str string) []byte {
	sha160 := sha1.Sum([]byte(str))
	return sha160[:]
}

func GenerateInfoHash(bytes []byte) []byte {
	sha160 := sha1.Sum(bytes)
	return sha160[:]
}
