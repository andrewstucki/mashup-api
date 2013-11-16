package helpers

import (
	"encoding/base64"
	"math/rand"
)

var RandomSource = rand.NewSource(1028890720402726901)

func SecureToken() string {
	dest := make([]byte, 16, 16)
	RandomBytes(dest)
	return base64.URLEncoding.EncodeToString(dest)[:22]
}

func RandomBytes(p []byte) (n int, err error) {
	todo := len(p)
	offset := 0
	for {
		val := int64(RandomSource.Int63())
		for i := 0; i < 8; i++ {
			p[offset] = byte(val)
			todo--
			if todo == 0 {
				return len(p), nil
			}
			offset++
			val >>= 8
		}
	}
	panic("unreachable")
}
