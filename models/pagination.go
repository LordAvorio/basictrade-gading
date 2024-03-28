package models

type Pagination struct {
	Data         interface{} `json:"data"`
	NextPage     int         `json:"next_page"`
	PreviousPage int         `json:"previous_page"`
	TotalPage    int         `json:"total_page"`
	CurrentPage  int         `json:"current_page"`
}
