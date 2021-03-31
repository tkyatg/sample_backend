package userqueryservice

import (
	"errors"

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
	var res []getUserListResult
	if result := da.db.Table("users").Select("user_uuid, display_name, gender, image_url, free_time, self_introduction, created_at, updated_at").Find(&res); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []getUserListResult{}, nil
		}
		return nil, result.Error
	}
	return res, nil
}

func (da *dataAccessor) getUserByID(userUUID string) (getUserByIDResult, error) {
	res := getUserByIDResult{}
	if result := da.db.Table("users").Select("user_uuid, display_name, gender, image_url, free_time, self_introduction, created_at, updated_at, deleted_at").Where("user_uuid = ?", userUUID).Find(&res); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return getUserByIDResult{}, nil
		}
		return getUserByIDResult{}, result.Error
	}
	return res, nil
}
