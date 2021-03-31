package userquery

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
	getUserByIDResult struct {
		UserUUID         string `gorm:"primary_key"`
		DisplayName      string `gorm:"column:display_name"`
		BirthDay         string `gorm:"column:birthday"`
		Gender           string `gorm:"column:gender"`
		ImageURL         string `gorm:"column:image_url"`
		FreeTime         string `gorm:"column:free_time"`
		SelfIntroduction string `gorm:"column:self_introduction"`
		CreatedAt        string `gorm:"column:created_at"`
		UpdatedAt        string `gorm:"column:updated_at"`
	}
	getUserListResult struct {
		UserUUID         string `gorm:"primary_key"`
		DisplayName      string `gorm:"column:display_name"`
		BirthDay         string `gorm:"column:birthday"`
		Gender           string `gorm:"column:gender"`
		ImageURL         string `gorm:"column:image_url"`
		FreeTime         string `gorm:"column:free_time"`
		SelfIntroduction string `gorm:"column:self_introduction"`
		CreatedAt        string `gorm:"column:created_at"`
		UpdatedAt        string `gorm:"column:updated_at"`
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
