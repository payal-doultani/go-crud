package response

type StudentResponse struct {
	ID       int    `json:"id" orm:"column(id)"`
	Name     string  `json:"name" orm:"column(name)"`
	Email     string  `json:"email" orm:"column(email)"`
}

type ErrorResponse struct {
    Message string `json:"message"`
}

type GetAllStudentsResponse struct {
	Students      []StudentResponse `json:"students"`
	PageNo      int             `json:"page_no"`
	PageSize    int             `json:"per_page"`
	TotalCount  int             `json:"total_count"`
	LastPage    int             `json:"last_page"`
	CurrentPage int             `json:"current_page"`
}
