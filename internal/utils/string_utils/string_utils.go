package stringutils

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func FirstLetterToLower(s string) string {

	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func HashPassword(str string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes)
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

