package domain

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type (
	Email                string
	EncryptedPassword    string
	Gender               string
	CreateUserAttributes struct {
		email             Email
		encryptedPassword EncryptedPassword
		gender            Gender
	}
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func NewCreateUserAttributes(
	email string,
	password string,
	gender string,
) (*CreateUserAttributes, error) {
	em, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	pass, err := newEncryptedPassword(password)
	if err != nil {
		return nil, err
	}
	gen, err := newGender(gender)
	if err != nil {
		return nil, err
	}

	return &CreateUserAttributes{
		email:             em,
		encryptedPassword: pass,
		gender:            gen,
	}, nil
}

func newEmail(value string) (Email, error) {
	if value == "" || !emailRegexp.MatchString(value) {
		return "", errors.New("emailの形式が正しくありません。")
	}
	return Email(value), nil
}

func newEncryptedPassword(value string) (EncryptedPassword, error) {
	if value == "" || len(value) < 6 {
		return "", errors.New("passwordは6文字以上を入力してください。")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 12)
	if err != nil {
		return "", err
	}
	return EncryptedPassword(hash), nil
}

func newGender(value string) (Gender, error) {
	gender := []string{"男", "女"}
	if value == "" {
		return "", errors.New("性別を選択してください。")
	}

	exists := false
	for _, gen := range gender {
		if gen == value {
			exists = true
			break
		}
	}
	if !exists {
		return "", errors.New("性別が不正な値です。")
	}
	return Gender(value), nil
}
