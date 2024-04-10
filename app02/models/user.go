package models

import "time"

type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Created_At time.Time
	Updated_At time.Time
}
