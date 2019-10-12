package autopilot

import (
	"addressbook/models"
)

type Proxy interface {
	Get(id string) ([]models.Contact, error)
	Upsert(contact models.Contact) error
}
