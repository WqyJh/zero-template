package rand_test

import (
	"testing"
	"zero-template/common/utils/rand"
)

func TestRandString(t *testing.T) {
	t.Log(rand.RandString(32))
}

func TestRandHex(t *testing.T) {
	t.Log(rand.RandHex(32))
}
