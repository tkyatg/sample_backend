package domain

type (
	User struct {
		UserUUID         string `gorm:"primary_key"`
		DisplayName      string `gorm:"column:display_name"`
		Email            string `gorm:"column:email"`
		Password         string `gorm:"column:password"`
		BirthDay         string `gorm:"column:birthday"`
		TelephoneNumber  string `gorm:"column:telephone_number"`
		Gender           string `gorm:"column:gender"`
		ImageURL         string `gorm:"column:image_url"`
		FreeTime         string `gorm:"column:free_time"`
		SelfIntroduction string `gorm:"column:self_introduction"`
		CreatedAt        string `gorm:"column:created_at"`
		UpdatedAt        string `gorm:"column:updated_at"`
		DeletedAt        string `gorm:"column:deleted_at"`
	}
)
