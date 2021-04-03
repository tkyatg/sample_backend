package usercommand

import (
	"fmt"

	"github.com/tkyatg/rental_redy_backend/services/user/domain"
)

type (
	usecase struct {
		userRepo domain.UserRepository
	}
	// Usecase interface
	Usecase interface {
		Create(req createRequest) error
	}
)

// NewUsecase function
func NewUsecase(userRepo domain.UserRepository) Usecase {
	return &usecase{userRepo}
}

func (t *usecase) Create(req createRequest) error {
	attr, err := domain.NewCreateUserAttributes(req.Email, req.Password, req.Gender)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return t.userRepo.Create(attr)
}
