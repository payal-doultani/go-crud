package request

type StudentRequest struct {
	Name  string `json:"name" validate:"required,gte=3,lte=10"`
	Email string `json:"email" validate:"required,email"`
}

type Req struct {
	PageNo   int    `json:"page" query:"page_no" validate:"min=1"`
	PageSize int    `json:"page_size" query:"page_size" validate:"min=1"`
	OrderBy  string `json:"sort_by" query:"order_by"`
	Order    string `json:"sort_order" query:"order" validate:"oneof=asc desc"`
	Filter   string `json:"filter" query:"filter"`
}
