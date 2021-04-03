package userquery

import (
	"database/sql"
)

type dataAccessor struct {
	db *sql.DB
}

// NewDataAccessor function
func NewDataAccessor(db *sql.DB) DataAccessor {
	return &dataAccessor{db}
}
func (t *dataAccessor) getUserList() ([]*getUserListResult, error) {
	sql := `
select user_uuid
     , coalesce(display_name, '') display_name
     , gender
     , coalesce(image_url, '') image_url
     , coalesce(free_time, '') free_time
     , coalesce(self_introduction, '') self_introduction
     , created_at
     , updated_at
  from users
 where deleted_at is null
   and display = true
`

	rows, err := t.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var res []*getUserListResult
	for rows.Next() {
		user := new(getUserListResult)
		rows.Scan(&user.UserUUID, &user.DisplayName, &user.Gender, &user.ImageURL, &user.FreeTime, &user.SelfIntroduction, &user.CreatedAt, &user.UpdatedAt)
		res = append(res, user)
	}

	return res, nil
}

func (t *dataAccessor) getUserByID(req getUserByIDRequest) (*getUserByIDResult, error) {
	res := new(getUserByIDResult)
	sql := `
select user_uuid
     , coalesce(display_name, '') display_name
     , gender
     , coalesce(image_url, '') image_url
     , coalesce(free_time, '') free_time
     , coalesce(self_introduction, '') self_introduction
     , created_at
     , updated_at
  from users
 where user_uuid = $1
   and display = true
   and deleted_at is null
`
	if err := t.db.QueryRow(sql, req.userUUID).Scan(&res.UserUUID, &res.DisplayName, &res.Gender, &res.ImageURL, &res.FreeTime, &res.SelfIntroduction, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}
	return res, nil
}
