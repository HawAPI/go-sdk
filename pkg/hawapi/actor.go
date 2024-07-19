package hawapi

import (
	"github.com/google/uuid"
)

const actorOrigin = "actors"

type Social struct {
	Social string `json:"social,omitempty"`
	Handle string `json:"handle,omitempty"`
	URL    string `json:"url,omitempty"`
}

type Actor struct {
	UUID        uuid.UUID `json:"uuid"`
	Href        string    `json:"href"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Nicknames   []string  `json:"nicknames,omitempty"`
	Socials     []Social  `json:"socials,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
	BirthDate   string    `json:"birth_date,omitempty"`
	DeathDate   string    `json:"death_date,omitempty"`
	Gender      int       `json:"gender,omitempty"`
	Seasons     []string  `json:"seasons,omitempty"`
	Awards      []string  `json:"awards,omitempty"`
	Character   string    `json:"character"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Images      []string  `json:"images,omitempty"`
	Sources     []string  `json:"sources,omitempty"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type CreateActor struct {
	FirstName   string   `json:"first_name,omitempty"`
	LastName    string   `json:"last_name,omitempty"`
	Nicknames   []string `json:"nicknames,omitempty"`
	Socials     []Social `json:"socials,omitempty"`
	Nationality string   `json:"nationality,omitempty"`
	BirthDate   string   `json:"birth_date,omitempty"`
	DeathDate   string   `json:"death_date,omitempty"`
	Gender      int      `json:"gender,omitempty"`
	Seasons     []string `json:"seasons,omitempty"`
	Awards      []string `json:"awards,omitempty"`
	Character   string   `json:"character,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

type PatchActor = CreateActor

type ActorResponse struct {
	BaseResponse
	Data Actor `json:"data"`
}

type ActorListResponse struct {
	BaseResponse
	Data []Actor `json:"data"`
}

// ListActors will get all actors
func (c *Client) ListActors(options ...QueryOptions) (ActorListResponse, error) {
	var actors []Actor
	var res ActorListResponse

	doRes, err := c.doGetRequest(actorOrigin, options, &actors)
	if err != nil {
		return res, err
	}

	res = ActorListResponse{
		BaseResponse: doRes,
		Data:         actors,
	}

	return res, nil
}

// FindActor will get a single item by uuid
func (c *Client) FindActor(id uuid.UUID) (ActorResponse, error) {
	var actor Actor
	var res ActorResponse

	doRes, err := c.doGetRequest(actorOrigin+"/"+id.String(), nil, &actor)
	if err != nil {
		return res, err
	}

	res = ActorResponse{
		BaseResponse: doRes,
		Data:         actor,
	}

	return res, nil
}

func (c *Client) RandomActor() (ActorResponse, error) {
	var actor Actor
	var res ActorResponse

	doRes, err := c.doGetRequest(actorOrigin+"/random", nil, &actor)
	if err != nil {
		return res, err
	}

	res = ActorResponse{
		BaseResponse: doRes,
		Data:         actor,
	}

	return res, nil
}

func (c *Client) CreateActor(s CreateActor) (Actor, error) {
	var actor Actor

	err := c.doPostRequest(actorOrigin, s, &actor)
	if err != nil {
		return actor, err
	}

	return actor, nil
}

func (c *Client) PatchActor(id uuid.UUID, p PatchActor) (Actor, error) {
	var actor Actor

	err := c.doPatchRequest(actorOrigin+"/"+id.String(), &p)
	if err != nil {
		return actor, err
	}

	res, err := c.FindActor(id)
	if err != nil {
		return actor, err
	}

	actor = res.Data
	return actor, nil
}

func (c *Client) DeleteActor(id uuid.UUID) error {
	return c.doDeleteRequest(actorOrigin + "/" + id.String())
}
