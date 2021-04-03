package domain

import (
	"database/sql"
)

type (
	userDataAccessor struct {
		db *sql.DB
	}
)

func NewUserDataAccessor(
	db *sql.DB,
) UserDataAccessor {
	return &userDataAccessor{db}
}

func (t *userDataAccessor) create(attr *CreateUserAttributes) error {
	sql := `insert into users ( email, password, gender, display)
			values($1, $2, $3, false)`
	if _, err := t.db.Exec(sql, attr.email, attr.encryptedPassword, attr.gender); err != nil {
		return err
	}
	return nil
}
