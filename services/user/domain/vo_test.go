package domain

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewCreateUserAttributes(t *testing.T) {
	t.Parallel()
}

func TestNewEmail(t *testing.T) {
	t.Parallel()
	t.Run("正常系", func(t *testing.T) {
		if res, err := newEmail("test@gmail.com"); string(res) != "test@gmail.com" || err != nil {
			t.Fatal(res, err)
		}
	})
	t.Run("空", func(t *testing.T) {
		if res, err := newEmail(""); res != "" || err.Error() != "emailの形式が正しくありません。" {
			t.Fatal(res, err)
		}
	})
	t.Run("email形式が正しくない", func(t *testing.T) {
		if res, err := newEmail("test@"); res != "" || err.Error() != "emailの形式が正しくありません。" {
			t.Fatal(res, err)
		}
	})
}

func TestNewEncryptedPassword(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		res, err := newEncryptedPassword("password")
		if err != nil {
			t.Fatal(res, err)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(string(res)), []byte("password")); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("空", func(t *testing.T) {
		if res, err := newEncryptedPassword(""); res != "" || err.Error() != "passwordは6文字以上を入力してください。" {
			t.Fatal(res, err)
		}
	})
	t.Run("長さ5", func(t *testing.T) {
		if res, err := newEncryptedPassword("paswd"); res != "" || err.Error() != "passwordは6文字以上を入力してください。" {
			t.Fatal(res, err)
		}
	})
	t.Run("長さ6", func(t *testing.T) {
		res, err := newEncryptedPassword("paswrd")
		if err != nil {
			t.Fatal(res, err)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(string(res)), []byte("paswrd")); err != nil {
			t.Fatal(err)
		}
	})
}

func TestNewGender(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		if res, err := newGender("男"); string(res) != "男" || err != nil {
			t.Fatal(res, err)
		}
	})
	t.Run("空", func(t *testing.T) {
		if res, err := newGender(""); res != "" || err.Error() != "性別を選択してください。" {
			t.Fatal(res, err)
		}
	})
	t.Run("長さ5", func(t *testing.T) {
		if res, err := newGender("あ"); res != "" || err.Error() != "性別が不正な値です。" {
			t.Fatal(res, err)
		}
	})
}
