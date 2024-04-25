package authdto

type LoginResponse struct {
	// ID       int `gorm:"type: varchar(255)" json:"id"`
	MSISDN   string `gorm:"type: varchar(255)" json:"msisdn"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	ID       int    `gorm:"type: varchar(255)" json:"id"`
	MSISDN   string `gorm:"type: varchar(255)" json:"msisdn"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	// Token    string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	MSISDN   string `gorm:"type: varchar(255)" json:"msisdn"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Username string `gorm:"type: varchar(255)" json:"username"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}
