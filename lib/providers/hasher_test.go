package providers

import (
	"testing"
)

func TestHasher(test *testing.T) {
	hasher := HasherImpl{
		iterations:  100,
		memory:      32 * 1024,
		parallelism: 80,
		keyLen:      320,
		saltLen:     100,
	}
	_, err := hasher.Hash("123")
	if err != nil {
		test.Error(err)
	}
}
