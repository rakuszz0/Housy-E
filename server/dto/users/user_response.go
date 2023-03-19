package usersdto

type UserResponse struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Username   string `json:"username"`
	ListAsRole string `json:"list_as_role"`
	Address    string `json:"address"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Image      string `json:"image"`
}
