package util

import "golang.org/x/crypto/bcrypt"

// HashPassword returns a bcrypt hash of the password.
// Credit goes to https://gowebexamples.com/password-hashing/.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success or an error on failure.
// Credit goes to https://gowebexamples.com/password-hashing/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
