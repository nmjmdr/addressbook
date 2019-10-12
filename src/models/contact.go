package models

import (
	"time"
)

type Contact struct {
	ID         string                 `json:"contact_id"`
	FirstName  string                 `json:"FirstName"`
	LastName   string                 `json:"LastName"`
	Type       string                 `json:"type"`
	Email      string                 `json:"Email"`
	Phone      string                 `json:"Phone"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	LeadSource string                 `json:"LeadSource"`
	Status     string                 `json:"Status"`
	Company    string                 `json:"Company"`
	Lists      []string               `json:"lists"`
	Custom     map[string]interface{} `json:"custom"`
}
