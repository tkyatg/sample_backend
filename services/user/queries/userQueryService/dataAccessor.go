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
