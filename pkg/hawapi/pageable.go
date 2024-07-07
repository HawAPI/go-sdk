package hawapi

type Pageable struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Sort  string `json:"sort"`
	Order string `json:"order"`
}
