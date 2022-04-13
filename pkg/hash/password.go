package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA1Hasher struct {
	salt string
}

func (h *SHA1Hasher) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
