package hawapi

import (
	"github.com/google/uuid"
)

const characterOrigin = "characters"

type Character struct {
	Uuid      uuid.UUID `json:"uuid"`
	Href      string    `json:"href"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nicknames []string  `json:"nicknames,omitempty"`
	Gender    int       `json:"gender"`
	Actor     string    `json:"actor"`
	BirthDate string    `json:"birth_date,omitempty"`
	DeathDate string    `json:"death_date,omitempty"`
	Thumbnail string    `json:"thumbnail"`
	Images    []string  `json:"images,omitempty"`
	Sources   []string  `json:"sources,omitempty"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CreateCharacter struct {
	Thumbnail string   `json:"thumbnail,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	Gender    int      `json:"gender,omitempty"`
	Actor     string   `json:"actor,omitempty"`
	Images    []string `json:"images,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	BirthDate string   `json:"birth_date,omitempty"`
	DeathDate string   `json:"death_date,omitempty"`
	Sources   []string `json:"sources,omitempty"`
}

type PatchCharacter = CreateCharacter

type CharacterResponse struct {
	BaseResponse
	Data Character `json:"data"`
}

type CharacterListResponse struct {
	BaseResponse
	Data []Character `json:"data"`
}

// ListCharacters will get all characters
func (c *Client) ListCharacters(options ...ListOptions) (CharacterListResponse, error) {
	var characters []Character
	var res CharacterListResponse

	opts := c.newListOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(characterOrigin, &opts, &characters)
	if err != nil {
		return res, err
	}

	res = CharacterListResponse{
		BaseResponse: doRes,
		Data:         characters,
	}

	return res, nil
}

// FindCharacter will get a single item by uuid
func (c *Client) FindCharacter(id uuid.UUID) (CharacterResponse, error) {
	var character Character
	var res CharacterResponse

	doRes, err := c.doGetRequest(characterOrigin+"/"+id.String(), nil, &character)
	if err != nil {
		return res, err
	}

	res = CharacterResponse{
		BaseResponse: doRes,
		Data:         character,
	}

	return res, nil
}

func (c *Client) RandomCharacter() (CharacterResponse, error) {
	var character Character
	var res CharacterResponse

	doRes, err := c.doGetRequest(characterOrigin+"/random", nil, character)
	if err != nil {
		return res, err
	}

	res = CharacterResponse{
		BaseResponse: doRes,
		Data:         character,
	}

	return res, nil
}

func (c *Client) CreateCharacter(s CreateCharacter) (Character, error) {
	var character Character

	err := c.doPostRequest(characterOrigin, s, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}

func (c *Client) PatchCharacter(id uuid.UUID, p PatchCharacter) (Character, error) {
	var character Character

	err := c.doPatchRequest(characterOrigin+"/"+id.String(), &p)
	if err != nil {
		return character, err
	}

	res, err := c.FindCharacter(id)
	if err != nil {
		return character, err
	}

	character = res.Data
	return character, nil
}

func (c *Client) DeleteCharacter(id uuid.UUID) error {
	return c.doDeleteRequest(characterOrigin + "/" + id.String())
}
