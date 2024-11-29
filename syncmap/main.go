package syncmap

import (
	"iter"
	"sync"
)

// SyncMap is a wrapper around sync.Map.
type SyncMap[K comparable, V any] struct {
	inner sync.Map
}

// Load returns the value stored in the map for a key, or nil if no value is present.
// The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	val, ok := m.inner.Load(key)
	if !ok {
		return *new(V), false
	}
	return val.(V), true
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	val, ok := m.inner.LoadOrStore(key, value)
	return val.(V), ok
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.inner.Store(key, value)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	val, ok := m.inner.LoadAndDelete(key)
	if !ok {
		return *new(V), false
	}
	return val.(V), true
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.inner.Delete(key)
}

func (m *SyncMap[K, V]) Range() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.inner.Range(func(key, value any) bool {
			return yield(key.(K), value.(V))
		})
	}
}

// Clear deletes all the entries, resulting in an empty Map.
func (m *SyncMap[K, V]) Clear() {
	m.inner.Clear()
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (m *SyncMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.inner.CompareAndSwap(key, old, new)
}

// CompareAndDelete deletes the entry for key if its value is equal to old. The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete returns false (even if the old value is the nil interface value).
func (m *SyncMap[K, V]) Swap(key K, old V) (deleted bool) {
	return m.inner.CompareAndDelete(key, old)
}
