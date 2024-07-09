package hawapi

import (
	"github.com/google/uuid"
)

const episodeOrigin = "episodes"

type Episode struct {
	Uuid        uuid.UUID `json:"uuid"`
	Href        string    `json:"href"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Duration    int64     `json:"duration"`
	Season      string    `json:"season"`
	EpisodeNum  byte      `json:"episode_num"`
	NextEpisode string    `json:"next_episode,omitempty"`
	PrevEpisode string    `json:"prev_episode,omitempty"`
	Thumbnail   string    `json:"thumbnail"`
	Images      []string  `json:"images,omitempty"`
	Sources     []string  `json:"sources,omitempty"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type CreateEpisode struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Language    string   `json:"language"`
	Duration    int64    `json:"duration"`
	Season      string   `json:"season"`
	EpisodeNum  byte     `json:"episode_num"`
	NextEpisode string   `json:"next_episode,omitempty"`
	PrevEpisode string   `json:"prev_episode,omitempty"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

type PatchEpisode = CreateEpisode

type EpisodeResponse struct {
	BaseResponse
	Data Episode `json:"data"`
}

type EpisodeListResponse struct {
	BaseResponse
	Data []Episode `json:"data"`
}

// ListEpisodes will get all episodes
func (c *Client) ListEpisodes(options ...QueryOptions) (EpisodeListResponse, error) {
	var episodes []Episode
	var res EpisodeListResponse

	opts := c.newQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(episodeOrigin, &opts, &episodes)
	if err != nil {
		return res, err
	}

	res = EpisodeListResponse{
		BaseResponse: doRes,
		Data:         episodes,
	}

	return res, nil
}

// FindEpisode will get a single item by uuid
func (c *Client) FindEpisode(id uuid.UUID) (EpisodeResponse, error) {
	var episode Episode
	var res EpisodeResponse

	opts := c.newQueryOptions()
	doRes, err := c.doGetRequest(episodeOrigin+"/"+id.String(), &opts, &episode)
	if err != nil {
		return res, err
	}

	res = EpisodeResponse{
		BaseResponse: doRes,
		Data:         episode,
	}

	return res, nil
}

func (c *Client) RandomEpisode() (EpisodeResponse, error) {
	var episode Episode
	var res EpisodeResponse

	opts := c.newQueryOptions()
	doRes, err := c.doGetRequest(episodeOrigin+"/random", &opts, &episode)
	if err != nil {
		return res, err
	}

	res = EpisodeResponse{
		BaseResponse: doRes,
		Data:         episode,
	}

	return res, nil
}

func (c *Client) CreateEpisode(s CreateEpisode) (Episode, error) {
	var episode Episode

	err := c.doPostRequest(episodeOrigin, s, &episode)
	if err != nil {
		return episode, err
	}

	return episode, nil
}

func (c *Client) PatchEpisode(id uuid.UUID, p PatchEpisode) (Episode, error) {
	var episode Episode

	err := c.doPatchRequest(episodeOrigin+"/"+id.String(), &p)
	if err != nil {
		return episode, err
	}

	res, err := c.FindEpisode(id)
	if err != nil {
		return episode, err
	}

	episode = res.Data
	return episode, nil
}

func (c *Client) DeleteEpisode(id uuid.UUID) error {
	return c.doDeleteRequest(episodeOrigin + "/" + id.String())
}
