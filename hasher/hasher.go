package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	Hash(text string) (string, error)
	Compare(hashed string, text string) bool
}

type bcryptHasher struct{}

func NewHasher() Hasher {
	return &bcryptHasher{}
}

func (h *bcryptHasher) Hash(text string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), 10)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (h *bcryptHasher) Compare(hashed string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	return err == nil
}
