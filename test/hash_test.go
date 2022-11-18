package benchmark

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"testing"

	"github.com/spaolacci/murmur3"
)

func benchmarkRun(h hash.Hash, i int, b *testing.B) {
	bs := make([]byte, i)
	_, err := rand.Read(bs)
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Write(bs)
		h.Sum(nil)
	}

}

func BenchmarkMD5_1k(b *testing.B) {
	benchmarkRun(md5.New(), 1024, b)
}

func BenchmarkMD5_10k(b *testing.B) {
	benchmarkRun(md5.New(), 10*1024, b)
}

func BenchmarkMD5_100k(b *testing.B) {
	benchmarkRun(md5.New(), 100*1024, b)
}

func BenchmarkSha1_1k(b *testing.B) {
	benchmarkRun(sha1.New(), 1024, b)
}

func BenchmarkSha1_10k(b *testing.B) {
	benchmarkRun(sha1.New(), 10*1024, b)
}

func BenchmarkSha1_100k(b *testing.B) {
	benchmarkRun(sha1.New(), 100*1024, b)
}

func BenchmarkSha256_1k(b *testing.B) {
	benchmarkRun(sha256.New(), 1024, b)
}

func BenchmarkSha256_10k(b *testing.B) {
	benchmarkRun(sha256.New(), 10*1024, b)
}

func BenchmarkSha256_100k(b *testing.B) {
	benchmarkRun(sha256.New(), 100*1024, b)
}

func BenchmarkMurmur3_32_1k(b *testing.B) {
	benchmarkRun(murmur3.New32(), 1024, b)
}

func BenchmarkMurmur3_32_10k(b *testing.B) {
	benchmarkRun(murmur3.New32(), 10*1024, b)
}

func BenchmarkMurmur3_32_100k(b *testing.B) {
	benchmarkRun(murmur3.New32(), 100*1024, b)
}

func BenchmarkMurmur3_64_1k(b *testing.B) {
	benchmarkRun(murmur3.New64(), 1024, b)
}

func BenchmarkMurmur3_64_10k(b *testing.B) {
	benchmarkRun(murmur3.New64(), 10*1024, b)
}

func BenchmarkMurmur3_64_100k(b *testing.B) {
	benchmarkRun(murmur3.New64(), 100*1024, b)
}

func BenchmarkMurmur3_128_1k(b *testing.B) {
	benchmarkRun(murmur3.New128(), 1024, b)
}

func BenchmarkMurmur3_128_10k(b *testing.B) {
	benchmarkRun(murmur3.New128(), 10*1024, b)
}

func BenchmarkMurmur3_128_100k(b *testing.B) {
	benchmarkRun(murmur3.New128(), 100*1024, b)
}
