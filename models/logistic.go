package models

import "time"

type Logistic struct {
	ID              int       `json:"id"`
	LogisticName    string    `json:"logistic_name" gorm:"type: varchar(255)"`
	Amount          int       `json:"amount"`
	DestinationName string    `json:"destination_name" gorm:"type: varchar(255)"`
	OriginName      string    `json:"origin_name" gorm:"type: varchar(255)"`
	Duration        string    `json:"duration" gorm:"type: varchar(255)"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
