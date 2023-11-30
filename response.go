package pkg

type ErrorResponse struct {
	Message string
}

type PaginatedResponse struct {
	Data  interface{} `json:"Data"`
	Total int64       `json:"Total"`
	Page  int         `json:"Page"`
	Limit int         `json:"Limit"`
}