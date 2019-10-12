package cache

// Cache - interface for the caching layer
type Cache interface {
	Add(key string, val interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
}
