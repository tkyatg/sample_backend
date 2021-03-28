package userqueryservice

type (
	usecase struct {
		da DataAccessor
	}
	// Usecase interface
	Usecase interface {
	}
	// DataAccessor interface
	DataAccessor interface {
	}
)

// NewUsecase function
func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da}
}
