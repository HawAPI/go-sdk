package hawapi

import "github.com/google/uuid"

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

// ListActors will get all actors
func (c *Client) ListActors() ([]Actor, error) {
	var actors []Actor

	err := c.get(actorOrigin, uuid.Nil, &actors)
	if err != nil {
		return actors, err
	}

	return actors, nil
}

// FindActor will get a single actor by UUID
func (c *Client) FindActor(id uuid.UUID) (Actor, error) {
	var actor Actor

	err := c.get(actorOrigin, id, &actor)
	if err != nil {
		return actor, err
	}

	return actor, nil
}

// RandomActor will get a single and random actor
func (c *Client) RandomActor() (Actor, error) {
	var actor Actor

	err := c.get(actorOrigin+"/random", uuid.Nil, &actor)
	if err != nil {
		return actor, err
	}

	return actor, nil
}
