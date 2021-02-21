package model

type Pagination struct {
	Total int64       `json:"total"`
	Skip  int64       `json:"skip"`
	Limit int64       `json:"limit"`
	Data  interface{} `json:"data"`
}
