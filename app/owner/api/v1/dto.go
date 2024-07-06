package v1

type ListOwnersResponse struct {
	Rows       []GetOwnerResponse `json:"rows"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
	Total      int                `json:"total"`
	TotalPages int                `json:"total_pages"`
}

type GetOwnerResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateOwnerBody struct {
	Name string `json:"name"`
}

type CreateOwnerResponse struct {
	Message string `json:"message"`
}

type UpdateOwnerBody struct {
	Name *string `json:"name"`
}

type UpdateOwnerResponse struct {
	Message string `json:"message"`
}

type DeleteOwnerResponse struct {
	Message string `json:"message"`
}
