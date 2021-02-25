package controller

type pagination struct {
	Total int64       `json:"total"`
	Skip  int64       `json:"skip"`
	Limit int64       `json:"limit"`
	Data  interface{} `json:"data"`
}

func createPagination(total int64, skip int64, limit int64, data interface{}) *pagination {
	return &pagination{
		Total: total,
		Skip:  skip,
		Limit: limit,
		Data:  data,
	}
}
