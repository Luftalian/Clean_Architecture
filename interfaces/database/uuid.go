package database

import "github.com/google/uuid"

type UUIDHandler interface {
	New() uuid.UUID
	Parse(s string) (uuid.UUID, error)
}
