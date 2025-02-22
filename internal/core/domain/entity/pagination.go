package entity

type Pagination struct {
	Page        int `json:"page"`
	TotalPage   int `json:"total_page"`
	TotalRecord int `json:"total_record"`
}

type PaginatedResponse[T any] struct {
	Items      []T        `json:"items"`
	Pagination Pagination `json:"pagination"`
}
