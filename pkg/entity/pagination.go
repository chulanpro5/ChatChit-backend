package entity

type Pagination struct {
	Page      uint        `json:"page"`
	Limit     uint        `json:"limit"`
	TotalPage uint        `json:"totalPage"`
	TotalRows uint64      `json:"totalRows"`
	Rows      interface{} `json:"rows"`
}
