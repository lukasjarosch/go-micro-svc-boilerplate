package datastore

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/lukasjarosch/go-micro-svc-boilerplate/config"
	"github.com/rs/xid"
)


type mysql struct {
	db *gorm.DB
}

func init() {
	Register(new(mysql))
}

// Init database
func (m *mysql) Init(config config.DatabaseConfiguration) error {

	db, err := gorm.Open(config.Dialect, config.Uri)
	if err != nil {
	    return err
	}

	// gorm does only log on ERRORs but we want to handle them ourselves
	db.LogMode(false)

	// auto-migrate database models
	db.AutoMigrate(&User{})

	m.db = db

	return nil
}

func (m *mysql) Close() error {
	return m.db.Close()
}

// CreateUser creates a new user record and returns it
func (m *mysql) CreateUser(user *User) error {

	user.ID = xid.New().String()
	user.Created = time.Now()
	user.Updated = user.Created
	user.Deleted = nil

	if err := m.db.Create(user); err != nil {
		return err.Error
	}

	return nil
}