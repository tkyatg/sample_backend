package userquery

type (
	usecase struct {
		da DataAccessor
	}
	// Usecase interface
	Usecase interface {
		getUserList() ([]*getUserListResult, error)
		getUserByID(req getUserByIDRequest) (*getUserByIDResult, error)
	}
	// DataAccessor interface
	DataAccessor interface {
		getUserList() ([]*getUserListResult, error)
		getUserByID(req getUserByIDRequest) (*getUserByIDResult, error)
	}
	getUserByIDResult struct {
		UserUUID         string
		DisplayName      string
		BirthDay         string
		Gender           string
		ImageURL         string
		FreeTime         string
		SelfIntroduction string
		CreatedAt        string
		UpdatedAt        string
	}
	getUserListResult struct {
		UserUUID         string
		DisplayName      string
		BirthDay         string
		Gender           string
		ImageURL         string
		FreeTime         string
		SelfIntroduction string
		CreatedAt        string
		UpdatedAt        string
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}

func (uc *usecase) getUserList() ([]*getUserListResult, error) {
	return uc.da.getUserList()
}

func (uc *usecase) getUserByID(req getUserByIDRequest) (*getUserByIDResult, error) {
	return uc.da.getUserByID(req)
}
