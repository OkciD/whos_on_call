package models

import "github.com/google/uuid"

type User struct {
	UID    uuid.UUID
	Name   string
	ApiKey string
}
