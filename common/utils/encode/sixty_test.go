package encode_test

import (
	"math"
	"math/rand"
	"testing"
	"time"
	"zero-template/common/utils/encode"

	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func TestEncodeSixty(t *testing.T) {
	n := rand.Uint64()
	encoded := encode.Encode10To62(n)
	decoded := encode.Decode62To10(encoded)
	t.Logf("n:%d encoded:%s decoded:%d", n, encoded, decoded)
	assert.Equal(t, n, decoded)
}

func TestEncodeSixtyMax(t *testing.T) {
	var n uint64 = math.MaxUint64
	encoded := encode.Encode10To62(n)
	decoded := encode.Decode62To10(encoded)
	t.Logf("n:%d encoded:%s decoded:%d", n, encoded, decoded)
	assert.Equal(t, n, decoded)
	assert.Equal(t, 22, len(encoded))
}

func BenchmarkSixty(b *testing.B) {
	n := rand.Uint64()
	for i := 0; i < b.N; i++ {
		encode.Encode10To62(n)
	}
}

func BenchmarkShortuuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encode.Shortuuid()
	}
}

func Benchmark3rdShortuuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shortuuid.New()
	}
}
