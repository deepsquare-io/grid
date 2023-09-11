package storage_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/sbatch-service/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InMemoryStorageSuite struct {
	suite.Suite
	storage storage.Storage
}

func (suite *InMemoryStorageSuite) BeforeTest(suiteName, testName string) {
	suite.storage = storage.NewInMemoryStorage()
}

func (suite *InMemoryStorageSuite) TestGet() {
	// Arrange: Set a key-value pair
	key := "testKey"
	value := "testValue"
	expiration := time.Second * 5
	err := suite.storage.Set(context.Background(), key, value, expiration)
	assert.NoError(suite.T(), err)

	// Act: Get the value
	retrievedValue, err := suite.storage.Get(context.Background(), key)

	// Assert: Verify the retrieved value and error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), value, retrievedValue)
}

func (suite *InMemoryStorageSuite) TestGetWithExpired() {
	// Arrange: Set an expired key-value pair
	key := "expiredKey"
	value := "expiredValue"
	expiration := time.Millisecond // Expire in 1 millisecond
	err := suite.storage.Set(context.Background(), key, value, expiration)
	assert.NoError(suite.T(), err)

	// Act: Wait for expiration and then try to Get the value
	time.Sleep(time.Millisecond * 10) // Wait for expiration
	retrievedValue, err := suite.storage.Get(context.Background(), key)

	// Assert: Verify the error is ErrNotFound
	assert.Equal(suite.T(), storage.ErrNotFound, err)
	assert.Empty(suite.T(), retrievedValue)
}

func (suite *InMemoryStorageSuite) TestGetWithNotFound() {
	// Act: Get a non-existent key
	key := "nonExistentKey"
	retrievedValue, err := suite.storage.Get(context.Background(), key)

	// Assert: Verify the error is ErrNotFound
	assert.Equal(suite.T(), storage.ErrNotFound, err)
	assert.Empty(suite.T(), retrievedValue)
}

func (suite *InMemoryStorageSuite) TestSet() {
	// Arrange: Set a key-value pair
	key := "testKey"
	value := "testValue"
	expiration := time.Second * 5

	// Act: Set the key-value pair
	err := suite.storage.Set(context.Background(), key, value, expiration)

	// Assert: Verify no error occurred
	assert.NoError(suite.T(), err)

	// Verify that we can retrieve the value
	retrievedValue, err := suite.storage.Get(context.Background(), key)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), value, retrievedValue)
}

func TestInMemoryStorageSuite(t *testing.T) {
	suite.Run(t, new(InMemoryStorageSuite))
}
