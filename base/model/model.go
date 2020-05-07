package model

import (
	"sync"

	"github.com/globalsign/mgo"
)

const (
	MongoRunMode     = "Strong"
	MongoPoolMax     = 4096
	MongoSyncTimeout = 5
)

type Model struct {
	mux        sync.RWMutex
	session    *mgo.Session
	collection *mgo.Collection

	config *Config
}
