package models

import "time"

type User struct {
	ID        int       `json:"id"`
	MSISDN    string    `json:"msisdn" gorm:"type: varchar(100)"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Username  string    `json:"username" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
