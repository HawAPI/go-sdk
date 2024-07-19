package hawapi

type DataCount struct {
	Actors      int `json:"actors"`
	Characters  int `json:"characters"`
	Episodes    int `json:"episodes"`
	Games       int `json:"games"`
	Locations   int `json:"locations"`
	Seasons     int `json:"seasons"`
	Soundtracks int `json:"soundtracks"`
}

type Overview struct {
	Uuid        string    `json:"uuid"`
	Href        string    `json:"href"`
	Sources     []string  `json:"sources"`
	Thumbnail   string    `json:"thumbnail"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Languages   []string  `json:"languages"`
	Creators    []string  `json:"creators"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DataCount   DataCount `json:"data_count"`
}

func (c *Client) Overview(options ...QueryOptions) (Overview, error) {
	var overview Overview

	_, err := c.doGetRequest("overview", options, &overview)
	if err != nil {
		return overview, err
	}

	return overview, nil
}
