package config

import (
	"fmt"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var Secret = []byte("Sudah izin M. Iqbal Ramadlani")

// VerifyToken verifies the given JWT token string and returns the token if valid.
func VerifyToken(tokenString string, secret []byte) (jwt.Token, error) {
	token, err := jwt.Parse([]byte(tokenString), jwt.WithValidate(true), jwt.WithVerify(jwa.HS256, secret))
	if err != nil {
		fmt.Printf("Gagal verifikasi token: %s\n", err)
		return nil, err
	}

	return token, nil
}
