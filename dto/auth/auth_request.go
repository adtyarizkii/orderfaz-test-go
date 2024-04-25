package authdto

type RegisterRequest struct {
	MSISDN   string `gorm:"type: varchar(255)" json:"msisdn" validate:"required"`
	Name     string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Username string `gorm:"type: varchar(255)" json:"username" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}

type LoginRequest struct {
	MSISDN   string `gorm:"type: varchar(255)" json:"msisdn" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
