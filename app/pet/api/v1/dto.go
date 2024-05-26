package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/entity"
)

type ListPetsResponse struct {
	Rows       []GetPetResponse `json:"rows"`
	Page       int              `json:"page"`
	PerPage    int              `json:"per_page"`
	Total      int              `json:"total"`
	TotalPages int              `json:"total_pages"`
}

type GetPetResponse struct {
	ID        string         `json:"id" gorm:"type:uuid;primary_key"`
	Name      string         `json:"name"`
	OwnerID   *string        `json:"owner_id"`
	Type      entity.PetType `json:"type"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

type CreatePetBody struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	OwnerID *string `json:"owner_id"`
}

type CreatePetResponse struct {
	Message string `json:"message"`
}

type UpdatePetBody struct {
	Name    *string `json:"name"`
	Type    *string `json:"type"`
	OwnerID *string `json:"owner_id"`
}

type UpdatePetResponse struct {
	Message string `json:"message"`
}

type DeletePetResponse struct {
	Message string `json:"message"`
}
