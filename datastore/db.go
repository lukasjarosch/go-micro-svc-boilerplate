package datastore

import (
	"github.com/lukasjarosch/go-micro-svc-boilerplate/config"
)


type DB interface {
	Init(config config.DatabaseConfiguration) error
	Close() error
	UserRepository
}

type UserRepository interface {
	CreateUser(user *User) error
}

func CreateUser(user *User) error {
	return db.CreateUser(user)
}


var (
	db DB
)

// Init initializes the database connection
func Init(config config.DatabaseConfiguration) error {
	return db.Init(config)
}

func Close() error {
	return db.Close()
}

// Register registers a database backend to use
func Register(backend DB) {
	db = backend
}



