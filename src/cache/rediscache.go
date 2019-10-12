package cache

import (
	"github.com/go-redis/redis"
	"addressbook/models"
	"encoding/json"
)
// implements cache using redis cache

/*
type Cache interface {
	Add(key string, val Contact) error
	Get(key string) (*Contact, error)
	Del(key string) error
}
*/


type RedisCache struct {
	client *redis.Client
}

func (r *RedisCache) Ping() error {
	_, err := r.client.Ping().Result()
	return err
}

func (r *RedisCache) Add(key string, val models.Contact) error{
	js, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return r.client.Set(key, js, 0).Err()
}

func (r *RedisCache)Get(key string) (*models.Contact, error) {
	val, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	contact := models.Contact{}
	err = json.Unmarshal([]byte(val), &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *RedisCache)Del(key string) error {
	return r.client.Del(key).Err()
}

func NewRedisCache(address string, password string) *RedisCache {
	r := &RedisCache{}
	options := redis.Options{
		Addr: address,
		DB: 0,
	}
	
	r.client = redis.NewClient(&options)
	return r
}

