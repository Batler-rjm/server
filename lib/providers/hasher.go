package providers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type HasherImpl struct {
	iterations  uint32
	memory      uint32
	parallelism uint8
	keyLen      uint32
	saltLen     uint32
}

func generateRandomBytes(len uint32) ([]byte, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (hasher HasherImpl) Hash(text string) (string, error) {
	salt, err := generateRandomBytes(hasher.saltLen)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(text), salt, hasher.iterations, hasher.memory, hasher.parallelism, hasher.keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		hasher.memory,
		hasher.iterations,
		hasher.parallelism,
		b64Salt,
		b64Hash,
	)
	return encodedHash, nil
}
