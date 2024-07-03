package hawapi

import "github.com/google/uuid"

const characterOrigin = "characters"

type Character struct {
	UUID      string   `json:"uuid"`
	Href      string   `json:"href"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Nicknames []string `json:"nicknames,omitempty"`
	BirthDate string   `json:"birth_date,omitempty"`
	DeathDate string   `json:"death_date,omitempty"`
	Gender    int      `json:"gender"`
	Thumbnail string   `json:"thumbnail,omitempty"`
	Actor     string   `json:"actor"`
	Images    []string `json:"images,omitempty"`
	Sources   []string `json:"sources,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// ListCharacters will get all characters
func (c *Client) ListCharacters() ([]Character, error) {
	var characters []Character

	err := c.get(characterOrigin, uuid.Nil, &characters)
	if err != nil {
		return characters, err
	}

	return characters, nil
}

// FindCharacter will get a single character by UUID
func (c *Client) FindCharacter(id uuid.UUID) (Character, error) {
	var character Character

	err := c.get(characterOrigin, id, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}

// RandomCharacter will get a single and random character
func (c *Client) RandomCharacter() (Character, error) {
	var character Character

	err := c.get(characterOrigin+"/random", uuid.Nil, &character)
	if err != nil {
		return character, err
	}

	return character, nil
}
