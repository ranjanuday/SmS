package auth

import "golang.org/x/crypto/bcrypt"

// =========================
// HASH PASSWORD
// =========================

func HashPassword(
	password string,
) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		14,
	)

	return string(bytes), err
}

// =========================
// COMPARE PASSWORDS
// =========================

func ComparePasswords(
	hashedPassword string,
	password []byte,
) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		password,
	)

	return err == nil
}