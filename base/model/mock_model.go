package model

import (
	"time"

	"github.com/globalsign/mgo"
)

var (
	_mockCollection       *_mockModelCollection
	mockCollection        = "mock"
	mockCollectionIndexes = []mgo.Index{
		// Compound indexes do not support the TTL property
		{
			Key:    []string{"name"},
			Unique: true,
		},
		// The TTL index is a single field index
		{
			Key:         []string{"created_at"},
			Unique:      false,
			ExpireAfter: 5 * time.Second,
		},
	}
)

type (
	_mockModelCollection struct{}
	MockModel struct {
		Name      string `bson:"name" json:"name"`
		BaseModel `bson:",inline"`
	}
)

func NewMockModel() *MockModel {
	return &MockModel{
		Name:      "mock",
		BaseModel: NewBaseModel(),
	}
}
