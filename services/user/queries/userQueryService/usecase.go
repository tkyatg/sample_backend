package userqueryservice

type (
	usecase struct {
		da DataAccessor
	}
	// Usecase interface
	Usecase interface {
		getUserList() ([]getUserListResult, error)
		getUserByID(userUUID string) (getUserByIDResult, error)
	}
	// DataAccessor interface
	DataAccessor interface {
		getUserList() ([]getUserListResult, error)
		getUserByID(userUUID string) (getUserByIDResult, error)
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}

func (uc *usecase) getUserList() ([]getUserListResult, error) {
	return uc.da.getUserList()
}
func (uc *usecase) getUserByID(userUUID string) (getUserByIDResult, error) {
	return uc.da.getUserByID(userUUID)
}
