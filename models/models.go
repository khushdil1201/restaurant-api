package models

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"-"`
    Role     string `json:"role"`
}

type MenuItem struct {
    gorm.Model
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}
