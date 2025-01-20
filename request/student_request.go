package request

type StudentRequest struct {
	Name  string `json:"name" validate:"required,gte=3,lte=10"`
	Email string `json:"email" validate:"required,email_regex"`
}
