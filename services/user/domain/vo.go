package domain

import (
	"errors"
	"regexp"
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
	return EncryptedPassword(value), nil
}

func newGender(value string) (Gender, error) {
	return Gender(value), nil
}
