package v1

type CreateUserInput struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Age      *uint64 `json:"age" validate:"required"`
}

type UpdateUserInput struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Age      *uint64 `json:"age" validate:"required"`
}
