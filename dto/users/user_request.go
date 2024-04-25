package usersdto

type CreateUserRequest struct {
	MSISDN   string `json:"msisdn" form:"msisdn" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type GetUserRequest struct {
	MSISDN   string `json:"msisdn" form:"msisdn" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
}

type UpdateUserRequest struct {
	MSISDN   string `json:"msisdn" form:"msisdn"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
