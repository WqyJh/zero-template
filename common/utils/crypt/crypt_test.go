package crypt_test

import (
	"testing"
	"zero-template/common/utils/crypt"
	"zero-template/common/utils/rand"

	"github.com/stretchr/testify/assert"
)

func encryptString(t *testing.T, s string, key string) string {
	encrypted, err := crypt.EncryptString(s, key)
	assert.NoError(t, err)
	return encrypted
}

func TestDecrypt(t *testing.T) {
	source := rand.RandString(32)
	key := rand.RandString(64)
	encrypted, err := crypt.Encrypt([]byte(source), key)
	assert.NoError(t, err)
	decrypted, err := crypt.Decrypt(encrypted, key)
	assert.NoError(t, err)
	assert.Equal(t, source, string(decrypted))
	assert.NotEqual(t, source, encrypted)
}

func TestDecode(t *testing.T) {
	type Nested struct {
		A string
		B []string
		C int
		D map[string][]string
	}
	type Config struct {
		A, B string
		C    int
		D    map[int]string
		E    map[string]Nested
		F    string
		G    []string
	}
	key := rand.RandString(64)
	expected := Config{
		A: rand.RandString(10),
		B: rand.RandString(128),
		C: 1,
		D: map[int]string{
			2: rand.RandString(20),
			8: rand.RandString(44),
		},
		E: map[string]Nested{
			"a": {
				A: rand.RandString(33),
				B: []string{rand.RandString(10), rand.RandString(20)},
				C: 1,
				D: map[string][]string{
					"b": {rand.RandString(10), rand.RandString(20)},
				},
			},
		},
		F: rand.RandString(20),
		G: []string{rand.RandString(10), rand.RandString(20)},
	}
	origin := Config{
		A: encryptString(t, expected.A, key),
		B: encryptString(t, expected.B, key),
		C: expected.C,
		D: map[int]string{
			2: encryptString(t, expected.D[2], key),
			8: encryptString(t, expected.D[8], key),
		},
		E: map[string]Nested{
			"a": {
				A: encryptString(t, expected.E["a"].A, key),
				B: []string{
					encryptString(t, expected.E["a"].B[0], key),
					encryptString(t, expected.E["a"].B[1], key),
				},
				C: expected.E["a"].C,
				D: map[string][]string{
					"b": {
						encryptString(t, expected.E["a"].D["b"][0], key),
						encryptString(t, expected.E["a"].D["b"][1], key),
					},
				},
			},
		},
		F: expected.F,
		G: []string{
			encryptString(t, expected.G[0], key),
			expected.G[1],
		},
	}

	result, err := crypt.Decode(origin, key)
	assert.NoError(t, err)
	assert.NotEqual(t, result, origin)
	assert.Equal(t, expected, result)
}
