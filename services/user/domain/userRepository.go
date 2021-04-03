package domain

type (
	userRepository struct {
		da UserDataAccessor
	}
	UserRepository interface {
		Create(attr *CreateUserAttributes) error
	}
	UserDataAccessor interface {
		create(attr *CreateUserAttributes) error
	}
)

func NewUserRepository(
	da UserDataAccessor,
) UserRepository {
	return &userRepository{da}
}

func (t *userRepository) Create(attr *CreateUserAttributes) error {
	return t.da.create(attr)
}
