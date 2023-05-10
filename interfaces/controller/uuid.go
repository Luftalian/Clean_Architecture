package controller

import "github.com/Luftalian/Clean_Architecture/domain"

type UUIDHandler interface {
	New() domain.UUID
	Parse(s string) (domain.UUID, error)
}
