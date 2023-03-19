package authdto

type SignUpRequest struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname" gorm:"type: varchar(255)"`
	Email      string `json:"email" gorm:"type: varchar(255)"`
	Password   string `json:"password" gorm:"type: varchar(255)"`
	Username   string `json:"username" gorm:"type: varchar(255)"`
	ListAsRole string `json:"list_as_role" gorm:"type: varchar(225)"`
	Gender     string `json:"gender" gorm:"type: varchar(225)"`
	Phone      string `json:"phone" gorm:"type: varchar(225)"`
	Address    string `json:"address" gorm:"type: varchar(225)"`
}

type SignInRequest struct {
	Username string `json:"username" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" gorm:"type: varchar(255)" validate:"required"`
	NewPassword string `json:"new_password" gorm:"type: varchar(255)" validate:"required"`
}
