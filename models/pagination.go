package models
type Pagination struct {
	Data       interface{} `json:"result_data"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Total      int         `json:"total"`
	PrevOffset int         `json:"previous_offset"`
	NextOffset int         `json:"next_offset"`
}
