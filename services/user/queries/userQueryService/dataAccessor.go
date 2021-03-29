package userqueryservice

import (
	"github.com/jinzhu/gorm"
)

type dataAccessor struct {
	db *gorm.DB
}

// NewDataAccessor function
func NewDataAccessor(db *gorm.DB) DataAccessor {
	return &dataAccessor{db}
}
func (da *dataAccessor) getUserList() ([]getUserListResult, error) {
	var users []getUserListResult
	if result := da.db.Table("users").Select("user_uuid, display_name, gender, image_url, free_time, self_introduction, created_at, updated_at").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (da *dataAccessor) getUserByID(userUUID string) (getUserByIDResult, error) {
	var user getUserByIDResult
	if result := da.db.Select("user_uuid, display_name, gender, image_url, free_time, self_introduction, created_at, updated_at, deleted_at").Where("user_uuid = ?", userUUID).Find(&user); result.Error != nil {
		return getUserByIDResult{}, result.Error
	}
	return user, nil
}
