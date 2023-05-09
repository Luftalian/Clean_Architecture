package infrastructure

import (
	"github.com/google/uuid"
)

type UUIDHandler struct{}

func (u *UUIDHandler) New() uuid.UUID {
	return uuid.New()
}

func (u *UUIDHandler) Parse(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}
