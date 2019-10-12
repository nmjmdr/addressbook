package store

import (
	"addressbook/models"
	"github.com/pkg/errors"
)

type ContactCache interface {
	Add(key string, contact models.Contact) error
	Get(key string) (*models.Contact, error)
	Del(key string) error
}

type AutopilotProxy interface {
	Get(idOrEmail string) ([]models.Contact, error)
	Upsert(contact models.Contact) error
}

type Store interface {
	Get(idOrEmail string) ([]models.Contact, error)
	Upsert(contact models.Contact) error
}

type store struct {
	cache ContactCache
	proxy AutopilotProxy
}

func (s *store) addToCache(contact models.Contact) error {
	err := s.cache.Add(contact.ID, contact)
	if err != nil {
		return err
	}
	err = s.cache.Add(contact.Email, contact)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) delFromCache(contact models.Contact) error {
	err := s.cache.Del(contact.ID)
	if err != nil {
		return err
	}
	err = s.cache.Del(contact.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *store) Get(idOrEmail string) ([]models.Contact, error) {
	val, err := s.cache.Get(idOrEmail)
	if err != nil {
		return []models.Contact{}, errors.Wrap(err,"Unable to check for cached value")
	}
	if val != nil {
		return []models.Contact{
			*val,
		}, nil
	}
	contacts, err := s.proxy.Get(idOrEmail)
	if err != nil {
		return []models.Contact{}, errors.Wrap(err, "Unable to fetch contact from Autopilot API")
	}

	if len(contacts) > 0 {
		// store with id as key, email as key
		err = s.addToCache(contacts[0])
		if err != nil {
			return []models.Contact{}, errors.Wrap(err, "Unable to store contact in cache")
		}
	}
	return contacts, nil
}

func (s *store) Upsert(contact models.Contact) error {
	err := s.proxy.Upsert(contact)
	if err != nil {
		return errors.Wrap(err, "Unable to update contact using Autopilot API")
	}
	err = s.delFromCache(contact)
	if err != nil {
		return errors.Wrap(err, "Unable to delete contact from cache")
	}
	return nil
}

func NewStore(cache ContactCache, proxy AutopilotProxy) Store {
	return &store {
		cache: cache,
		proxy: proxy,
	}
}
