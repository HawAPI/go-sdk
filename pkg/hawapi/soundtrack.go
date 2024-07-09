package hawapi

import (
	"github.com/google/uuid"
)

const soundtrackOrigin = "soundtracks"

type Soundtrack struct {
	UUID        uuid.UUID `json:"uuid"`
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	Duration    int64     `json:"duration"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album,omitempty"`
	ReleaseDate string    `json:"release_date"`
	Urls        []string  `json:"urls"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Images      []string  `json:"images,omitempty"`
	Sources     []string  `json:"sources,omitempty"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type CreateSoundtrack struct {
	Name        string   `json:"name,omitempty"`
	Duration    int64    `json:"duration,omitempty"`
	Artist      string   `json:"artist,omitempty"`
	Album       string   `json:"album,omitempty"`
	ReleaseDate string   `json:"release_date"`
	Urls        []string `json:"urls,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

type PatchSoundtrack = CreateSoundtrack

type SoundtrackResponse struct {
	BaseResponse
	Data Soundtrack `json:"data"`
}

type SoundtrackListResponse struct {
	BaseResponse
	Data []Soundtrack `json:"data"`
}

// ListSoundtracks will get all soundtracks
func (c *Client) ListSoundtracks(options ...QueryOptions) (SoundtrackListResponse, error) {
	var soundtracks []Soundtrack
	var res SoundtrackListResponse

	opts := c.newQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(soundtrackOrigin, &opts, &soundtracks)
	if err != nil {
		return res, err
	}

	res = SoundtrackListResponse{
		BaseResponse: doRes,
		Data:         soundtracks,
	}

	return res, nil
}

// FindSoundtrack will get a single item by uuid
func (c *Client) FindSoundtrack(id uuid.UUID) (SoundtrackResponse, error) {
	var soundtrack Soundtrack
	var res SoundtrackResponse

	doRes, err := c.doGetRequest(soundtrackOrigin+"/"+id.String(), nil, &soundtrack)
	if err != nil {
		return res, err
	}

	res = SoundtrackResponse{
		BaseResponse: doRes,
		Data:         soundtrack,
	}

	return res, nil
}

func (c *Client) RandomSoundtrack() (SoundtrackResponse, error) {
	var soundtrack Soundtrack
	var res SoundtrackResponse

	doRes, err := c.doGetRequest(soundtrackOrigin+"/random", nil, &soundtrack)
	if err != nil {
		return res, err
	}

	res = SoundtrackResponse{
		BaseResponse: doRes,
		Data:         soundtrack,
	}

	return res, nil
}

func (c *Client) CreateSoundtrack(s CreateSoundtrack) (Soundtrack, error) {
	var soundtrack Soundtrack

	err := c.doPostRequest(soundtrackOrigin, s, &soundtrack)
	if err != nil {
		return soundtrack, err
	}

	return soundtrack, nil
}

func (c *Client) PatchSoundtrack(id uuid.UUID, p PatchSoundtrack) (Soundtrack, error) {
	var soundtrack Soundtrack

	err := c.doPatchRequest(soundtrackOrigin+"/"+id.String(), &p)
	if err != nil {
		return soundtrack, err
	}

	res, err := c.FindSoundtrack(id)
	if err != nil {
		return soundtrack, err
	}

	soundtrack = res.Data
	return soundtrack, nil
}

func (c *Client) DeleteSoundtrack(id uuid.UUID) error {
	return c.doDeleteRequest(soundtrackOrigin + "/" + id.String())
}
