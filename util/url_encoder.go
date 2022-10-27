package util

import "github.com/spaolacci/murmur3"

// encode origin url to tiny url
// @param origin url
//
// @return string tiny url
func EncodeUrlByHash(origin string) string {
	hasher := murmur3.New32()
	hasher.Write([]byte(origin))

	return Base10ToBase62(uint64(hasher.Sum32()))
}
