package hasher_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/hasher"
	"github.com/stretchr/testify/assert"
)

func TestBcryptHasher_Hash(t *testing.T) {
	hash := hasher.NewHasher()
	text := "alice123"

	hashed, err := hash.Hash(text)

	assert.NotEqual(t, text, hashed)
	assert.NoError(t, err)
}

func TestBcryptHasher_Compare(t *testing.T) {
	hash := hasher.NewHasher()
	hashed := "$2a$10$5kC7huxBzWVu11sVtZdv5OdK7M7EhwPCiYO9EcDyKaIB.ObfmoCb6"
	text := "alice123"

	isValid := hash.Compare(hashed, text)

	assert.True(t, isValid)
}
