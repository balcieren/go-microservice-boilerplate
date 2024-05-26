package entity

import "github.com/google/uuid"

type PetType string

const (
	Dog    PetType = "dog"
	Cat    PetType = "cat"
	Bird   PetType = "bird"
	Fish   PetType = "fish"
	Rabbit PetType = "rabbit"
)

func (p PetType) String() string {
	return string(p)
}

func (p PetType) IsValid() bool {
	switch p {
	case Dog, Cat, Bird, Fish, Rabbit:
		return true
	}
	return false
}

type Pet struct {
	Base
	OwnerID *uuid.UUID `json:"owner_id"`
	Name    string     `json:"name"`
	Type    PetType    `json:"type"`
}
