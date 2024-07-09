package hawapi

import (
	"github.com/google/uuid"
)

const seasonOrigin = "seasons"

type Season struct {
	Uuid          uuid.UUID `json:"uuid"`
	Href          string    `json:"href"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Language      string    `json:"language"`
	Genres        []string  `json:"genres,omitempty"`
	Episodes      []string  `json:"episodes,omitempty"`
	Trailers      []string  `json:"trailers,omitempty"`
	Budget        int       `json:"budget"`
	DurationTotal int64     `json:"duration_total"`
	SeasonNum     byte      `json:"season_num"`
	ReleaseDate   string    `json:"release_date"`
	NextSeason    string    `json:"next_season,omitempty"`
	PrevSeason    string    `json:"prev_season,omitempty"`
	Thumbnail     string    `json:"thumbnail,omitempty"`
	Images        []string  `json:"images,omitempty"`
	Sources       []string  `json:"sources,omitempty"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

type CreateSeason struct {
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	Language      string   `json:"language,omitempty"`
	Genres        []string `json:"genres,omitempty"`
	Episodes      []string `json:"episodes,omitempty"`
	Trailers      []string `json:"trailers,omitempty"`
	Budget        int      `json:"budget,omitempty"`
	DurationTotal int64    `json:"duration_total,omitempty"`
	SeasonNum     byte     `json:"season_num,omitempty"`
	ReleaseDate   string   `json:"release_date,omitempty"`
	NextSeason    string   `json:"next_season,omitempty"`
	PrevSeason    string   `json:"prev_season,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty"`
	Images        []string `json:"images,omitempty"`
	Sources       []string `json:"sources,omitempty"`
}

type PatchSeason = CreateSeason

type SeasonResponse struct {
	BaseResponse
	Data Season `json:"data"`
}

type SeasonListResponse struct {
	BaseResponse
	Data []Season `json:"data"`
}

// ListSeasons will get all seasons
func (c *Client) ListSeasons(options ...QueryOptions) (SeasonListResponse, error) {
	var seasons []Season
	var res SeasonListResponse

	opts := c.newQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(seasonOrigin, &opts, &seasons)
	if err != nil {
		return res, err
	}

	res = SeasonListResponse{
		BaseResponse: doRes,
		Data:         seasons,
	}

	return res, nil
}

// FindSeason will get a single item by uuid
func (c *Client) FindSeason(id uuid.UUID) (SeasonResponse, error) {
	var season Season
	var res SeasonResponse

	doRes, err := c.doGetRequest(seasonOrigin+"/"+id.String(), nil, &season)
	if err != nil {
		return res, err
	}

	res = SeasonResponse{
		BaseResponse: doRes,
		Data:         season,
	}

	return res, nil
}

func (c *Client) RandomSeason() (SeasonResponse, error) {
	var season Season
	var res SeasonResponse

	doRes, err := c.doGetRequest(seasonOrigin+"/random", nil, &season)
	if err != nil {
		return res, err
	}

	res = SeasonResponse{
		BaseResponse: doRes,
		Data:         season,
	}

	return res, nil
}

func (c *Client) CreateSeason(s CreateSeason) (Season, error) {
	var season Season

	err := c.doPostRequest(seasonOrigin, s, &season)
	if err != nil {
		return season, err
	}

	return season, nil
}

func (c *Client) PatchSeason(id uuid.UUID, p PatchSeason) (Season, error) {
	var season Season

	err := c.doPatchRequest(seasonOrigin+"/"+id.String(), &p)
	if err != nil {
		return season, err
	}

	res, err := c.FindSeason(id)
	if err != nil {
		return season, err
	}

	season = res.Data
	return season, nil
}

func (c *Client) DeleteSeason(id uuid.UUID) error {
	return c.doDeleteRequest(seasonOrigin + "/" + id.String())
}
