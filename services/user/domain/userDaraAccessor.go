package domain

import (
	"database/sql"
	"errors"
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
	emailCheckSql := `
insert into
 users ( email
       , password
       , gender
       , display )
values ( $1
       , $2
       , $3
       , false)
`
	alreadyExists := true
	if err := t.db.QueryRow(emailCheckSql, attr.email).Scan(&alreadyExists); err != nil {
		return err
	}
	if alreadyExists {
		return errors.New("入力されたemailはすでに使われています。")
	}
	sql := `
insert into
 users ( email
       , password
       , gender
       , display )
values ( $1
	   , $2
	   , $3
	   , false)`
	if _, err := t.db.Exec(sql, attr.email, attr.encryptedPassword, attr.gender); err != nil {
		return err
	}
	return nil
}
