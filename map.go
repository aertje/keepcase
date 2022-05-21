package keepcase

import (
	"strings"
	"sync"
)

type Map[T any] struct {
	keys map[string]string
	m    map[string]T
	lock *sync.RWMutex
}

func NewMap[T any]() *Map[T] {
	return &Map[T]{
		keys: make(map[string]string),
		m:    make(map[string]T),
		lock: &sync.RWMutex{},
	}
}

// SetCaseRespect sets the value in the map for the given key, respecting the casing of a key
// if it already exists.
func (m Map[T]) SetCaseRespect(key string, value T) {
	canonicalKey := strings.ToLower(key)

	m.lock.Lock()
	defer m.lock.Unlock()

	if existingKey, ok := m.keys[canonicalKey]; ok {
		m.m[existingKey] = value
	} else {
		m.m[key] = value
		m.keys[canonicalKey] = key
	}

}

// SetCaseOverride sets the value in the map for the given key, overriding any existing key's
// casing.
func (m Map[T]) SetCaseOverride(key string, value T) {
	canonicalKey := strings.ToLower(key)

	m.lock.Lock()
	defer m.lock.Unlock()

	if existingKey, ok := m.keys[canonicalKey]; ok {
		delete(m.m, existingKey)
	}

	m.m[key] = value
	m.keys[canonicalKey] = key
}

// GetCaseInsensitive returns the value for the given key, ignoring the casing of the key
func (m Map[T]) GetCaseInsensitive(key string) (T, bool) {
	canonicalKey := strings.ToLower(key)

	m.lock.RLock()
	defer m.lock.RUnlock()

	if existingKey, ok := m.keys[canonicalKey]; ok {
		return m.m[existingKey], true
	}

	return *new(T), false
}

// GetCaseSensitive returns the value for the given key, respecting the casing of the key
func (m Map[T]) GetCaseSensitive(key string) (T, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok := m.m[key]
	return value, ok
}

// Get is equivalent to GetCaseInsensitive
func (m Map[T]) Get(key string) (T, bool) {
	return m.GetCaseInsensitive(key)
}

// Set is equivalent to SetCaseRespect
func (m Map[T]) Set(key string, value T) {
	m.SetCaseRespect(key, value)
}

func (m Map[T]) AsMap() map[string]T {
	return m.m
}

func (m Map[T]) Len() int {
	return len(m.keys)
}
