package misc

func Xor(a, b []byte) []byte {
	minLen := len(a)
	if minLen > len(b) {
		minLen = len(b)
	}
	dst := make([]byte, minLen)
	for i := 0; i < minLen; i++ {
		dst[i] = a[i] ^ b[i]
	}

	return dst
}
