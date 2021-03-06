package model

import (
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

const (
	MongoRunMode     = "Strong"
	MongoPoolMax     = 4096
	MongoSyncTimeout = 5

	DefaultMaxItems = 1000
	DefaultMinItems = 1
)


type Model struct {
	mux        sync.RWMutex
	session    *mgo.Session
	collection *mgo.Collection

	config  *Config
	logger  *logrus.Logger
	indexes map[string]bool
}

func NewModel(config *Config, logger *logrus.Logger) *Model {
	dsn := config.DSN()

	session, err := mgo.Dial(dsn)
	if err != nil {
		logger.Panic(err.Error())
	}

	if err := session.Ping(); err != nil {
		logger.Panic(err.Error())
	}

	// set session mode
	switch config.Mode {
	case MongoRunMode:
		session.SetMode(mgo.Strong, true)
	case "Monotonic":
		session.SetMode(mgo.Monotonic, true)
	case "Eventual":
		session.SetMode(mgo.Eventual, true)
	default:
		session.SetMode(mgo.Strong, true)
	}

	// set session safe
	session.SetSafe(&mgo.Safe{
		W:        1,
		WTimeout: 200,
	})

	// set pool size
	if config.Pool > 0 {
		if config.Pool > MongoPoolMax {
			config.Pool = MongoPoolMax
		}

		session.SetPoolLimit(config.Pool)
	}

	// set op response timeout
	if config.Timeout == 0 {
		config.Timeout = MongoSyncTimeout
	}
	session.SetSyncTimeout(time.Duration(config.Timeout) * time.Second)

	// panic as early as possible
	if err := session.Ping(); err != nil {
		panic(err.Error())
	}

	return &Model{
		session: session,
		config:  config,
		logger:  logger,
		indexes: make(map[string]bool),
	}
}

func (model *Model) Use(database string) *Model {
	model.config.Database = database

	return model
}

func (model *Model) Copy() *Model {
	return &Model{
		session: model.session.Copy(),
		config:  model.config.Copy(),
		logger:  model.logger,
	}
}

func (model *Model) C(name string) *Model {
	copiedDB := model.Copy()
	copiedDB.collection = copiedDB.session.DB(model.Database()).C(name)

	return copiedDB
}

func (model *Model) Config() *Config {
	return model.config
}

func (model *Model) Database() string {
	return model.config.Database
}

func (model *Model) Session() *mgo.Session {
	return model.session
}

func (model *Model) Collection() *mgo.Collection {
	return model.collection
}

func (model *Model) Query(collectionName string, collectionIndexes []mgo.Index, query func(*mgo.Collection)) {
	copiedDB := model.C(collectionName)
	defer copiedDB.Close()

	copiedCollection := copiedDB.Collection()

	if !model.indexes[collectionName] {
		model.mux.Lock()
		if !model.indexes[collectionName] {
			for _, index := range collectionIndexes {
				if err := copiedCollection.EnsureIndex(index); err != nil {
					model.indexes[collectionName] = false

					model.logger.Printf("Ensure index of %s (%#v) : %v", collectionName, index, err)
				}
			}

			model.indexes[collectionName] = true
		}
		model.mux.Unlock()
	}

	query(copiedCollection)
}

func (model *Model) Close() {
	model.session.Close()
}
