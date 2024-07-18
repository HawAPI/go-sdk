package hawapi

// Quota represents the quota status
type Quota struct {
	Remaining int `json:"remaining,omitempty"`
}

// HeaderResponse represents the formatted header response from a request
type HeaderResponse struct {
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
	PageTotal int    `json:"page_total,omitempty"`
	ItemSize  int    `json:"item_size,omitempty"`
	NextPage  int    `json:"next_page,omitempty"`
	PrevPage  int    `json:"prev_page,omitempty"`
	Language  string `json:"language,omitempty"`
	Quota     Quota  `json:"quota"`
	Etag      string `json:"etag"`
	Length    int    `json:"length"`
}

// BaseResponse represents all required response fields
type BaseResponse struct {
	HeaderResponse
	Cached bool `json:"cached,omitempty"`
	Status int  `json:"status"`
}
