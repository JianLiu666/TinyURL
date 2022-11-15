package tools

import (
	"math"
	"strings"
)

var base62 string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Convert base10 to base62 format
// @param num base10 format
//
// @return string base62 format
func Base10ToBase62(num uint64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, base62[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

// Convert base62 to base10
// @param str base62 format
//
// @return uint64 base10 format
func Base62ToBase10(str string) uint64 {
	var num uint64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(base62, str[i])
		num += uint64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

// reverse byte array (in-place)
// @param b target
func reverse(b []byte) {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
}
