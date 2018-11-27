package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/rene00/khaos/pkg/setting"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"strings"
)

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

// Credit to https://gist.github.com/tscholl2/dc7dc15dc132ea70a98e8542fefffa28
func derivePbkdf2Key(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

// Credit to https://gist.github.com/tscholl2/dc7dc15dc132ea70a98e8542fefffa28
func EncryptString(plaintext string) string {
	key, salt := derivePbkdf2Key(setting.AppSetting.Secret, nil)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

// Credit to https://gist.github.com/tscholl2/dc7dc15dc132ea70a98e8542fefffa28
func DecryptString(ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := derivePbkdf2Key(setting.AppSetting.Secret, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}
