package data

import (
	"github.com/devgek/webskeleton/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

//MockedDatastore ...
type MockedDatastore struct {
	mock.Mock
}

//GetDB ...
func (m *MockedDatastore) GetDB() *gorm.DB {
	return nil
}

//GetUser ...
func (m *MockedDatastore) GetUser(username string) (*models.User, error) {

	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

//CreateEntity ...
func (m *MockedDatastore) CreateEntity(entity interface{}) error {
	args := m.Called(entity)

	return args.Error(0)
}

//SaveEntity ...
func (m *MockedDatastore) SaveEntity(entity interface{}) error {
	args := m.Called(entity)

	return args.Error(0)
}

//DeleteEntityByID ...
func (m *MockedDatastore) DeleteEntityByID(entity interface{}, id uint) error {

	args := m.Called(entity, id)
	return args.Error(0)
}

//GetAllEntities ...
func (m *MockedDatastore) GetAllEntities(entitySlice interface{}) error {
	args := m.Called(entitySlice)

	return args.Error(0)
}

//GetOneEntityBy ...
func (m *MockedDatastore) GetOneEntityBy(entity interface{}, key string, val interface{}) error {
	args := m.Called(entity, key, val)

	return args.Error(0)
}

//GetEntityByID ...
func (m *MockedDatastore) GetEntityByID(entity interface{}, id uint) error {
	args := m.Called(entity, id)

	return args.Error(0)
}
