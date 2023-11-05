package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}
