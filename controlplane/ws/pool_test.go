package ws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextKey(t *testing.T) {
	assert.Equal(t, string(PoolContextKey), "pool", "should be 'pool'")
	assert.Equal(t, contextKey("pool"), PoolContextKey, "should be 'pool'")
}

func TestNewConnectionPool(t *testing.T) {
	pool := NewConnectionPool()
	assert.NotNil(t, pool, "should not be nil")
	assert.NotNil(t, pool.connections, "should not be nil")
	assert.Equal(t, len(pool.connections), 0, "should be 0")
}

func TestAdd(t *testing.T) {
	pool := NewConnectionPool()
	assert.NotNil(t, pool, "should not be nil")
	assert.NotNil(t, pool.connections, "should not be nil")
	assert.Equal(t, len(pool.connections), 0, "should be 0")
	assert.True(t, pool.Add(nil), "should be true")
	assert.Equal(t, len(pool.connections), 1, "should be 1")
}

func TestRemove(t *testing.T) {
	pool := NewConnectionPool()
	assert.NotNil(t, pool, "should not be nil")
	assert.NotNil(t, pool.connections, "should not be nil")
	assert.Equal(t, len(pool.connections), 0, "should be 0")
	assert.True(t, pool.Add(nil), "should be true")
	assert.Equal(t, len(pool.connections), 1, "should be 1")
	assert.True(t, pool.Remove(nil), "should be true")
	assert.Equal(t, len(pool.connections), 0, "should be 0")
}
