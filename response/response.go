package response

type Error struct {
	Message string
}

type Paginated struct {
	Data  interface{} `json:"Data"`
	Total int64       `json:"Total"`
	Page  int         `json:"Page"`
	Limit int         `json:"Limit"`
}