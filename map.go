package keepcase

import (
	"strings"
	"sync"
)

type Map[T any] struct {
	keys    map[string]string
	backing map[string]T
	lock    *sync.RWMutex
}

func NewMap[T any](backing map[string]T) *Map[T] {
	if backing == nil {
		backing = make(map[string]T)
	}

	m := &Map[T]{
		keys:    make(map[string]string),
		backing: backing,
		lock:    &sync.RWMutex{},
	}

	for k := range backing {
		m.keys[strings.ToLower(k)] = k
	}

	return m
}

// SetCaseRespect sets the value in the map for the given key, respecting the casing of a key
// if it already exists.
func (m Map[T]) SetCaseRespect(key string, value T) {
	canonicalKey := strings.ToLower(key)

	m.lock.Lock()
	defer m.lock.Unlock()

	if existingKey, ok := m.keys[canonicalKey]; ok {
		m.backing[existingKey] = value
	} else {
		m.backing[key] = value
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
		delete(m.backing, existingKey)
	}

	m.backing[key] = value
	m.keys[canonicalKey] = key
}

// GetCaseInsensitive returns the value for the given key, ignoring the casing of the key
func (m Map[T]) GetCaseInsensitive(key string) (T, bool) {
	canonicalKey := strings.ToLower(key)

	m.lock.RLock()
	defer m.lock.RUnlock()

	if existingKey, ok := m.keys[canonicalKey]; ok {
		return m.backing[existingKey], true
	}

	return *new(T), false
}

// GetCaseSensitive returns the value for the given key, respecting the casing of the key
func (m Map[T]) GetCaseSensitive(key string) (T, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok := m.backing[key]
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

func (m Map[T]) SetCollection(input map[string]T) {
	for k, v := range input {
		m.SetCaseRespect(k, v)
	}
}

func (m Map[T]) GetBacking() map[string]T {
	return m.backing
}

func (m Map[T]) Len() int {
	return len(m.keys)
}
