package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"

	"github.com/joshcwainwright/coffer/internal/config"
)

type AESGCM struct {
	aead cipher.AEAD
}

const (
	KeySize   = 32
	NonceSize = 12
	TagSize   = 16
	Overhead  = NonceSize + TagSize
)

var ErrDecrypt = errors.New("ciphertext is corrupt or was encrypted under a different master key")

func New(key config.MasterKey) (*AESGCM, error) {
	k := key.Value()
	if len(k) != KeySize {
		return nil, fmt.Errorf("master key must be %d bytes, got %d", KeySize, len(k))
	}

	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, fmt.Errorf("master key: %w", err)
	}

	aead, err := cipher.NewGCMWithRandomNonce(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}

	return &AESGCM{aead: aead}, nil
}

func (c *AESGCM) Encrypt(plaintext []byte) ([]byte, error) {
	return c.aead.Seal(nil, nil, plaintext, nil), nil
}

func (c *AESGCM) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := c.aead.Open(nil, nil, ciphertext, nil)
	if err != nil {
		return nil, ErrDecrypt
	}

	return plaintext, nil
}
