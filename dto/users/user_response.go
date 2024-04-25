package usersdto

type UserResponse struct {
	ID       string `json:"id"`
	MSISDN   string `json:"msisdn" form:"msisdn" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
