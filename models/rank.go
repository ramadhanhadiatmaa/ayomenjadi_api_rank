package models

import "gorm.io/gorm"

type Rank struct {
	gorm.Model
	Username string `json:"username"`
	Result   int    `json:"result"`
	Tryout   string `json:"tryout"`
}
