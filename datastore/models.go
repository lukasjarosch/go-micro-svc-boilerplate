package datastore

import (
	"time"
)

type User struct {
	ID string
	Name string
	Email string `gorm:"type:varchar(100);"`
	Created time.Time
	Updated time.Time
	Deleted *time.Time
}
