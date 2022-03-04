package request

type AlumnusNameRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=42"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
