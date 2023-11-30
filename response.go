package pkg

type ResponseError struct {
	Message string
}

type ResponsePaginated struct {
	Data  interface{} `json:"Data"`
	Total int64       `json:"Total"`
	Page  int         `json:"Page"`
	Limit int         `json:"Limit"`
}