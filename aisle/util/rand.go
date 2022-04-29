/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/4/29 17:24
 */

package util

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GenerateRandomBytes(length int) []byte {
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < 256 {
			b[i] = byte(idx)
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
