package hawapi

import (
	"github.com/google/uuid"
)

const gameOrigin = "games"

type Game struct {
	Uuid        string   `json:"uuid"`
	Href        string   `json:"href"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Playtime    int64    `json:"playtime"`
	Language    string   `json:"language"`
	Platforms   []string `json:"platforms,omitempty"`
	Stores      []string `json:"stores,omitempty"`
	Modes       []string `json:"modes,omitempty"`
	Genres      []string `json:"genres,omitempty"`
	Publishers  []string `json:"publishers,omitempty"`
	Developers  []string `json:"developers,omitempty"`
	Website     string   `json:"website"`
	Tags        []string `json:"tags,omitempty"`
	Trailer     string   `json:"trailer"`
	AgeRating   string   `json:"age_rating"`
	ReleaseDate string   `json:"release_date"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type CreateGame struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Playtime    int64    `json:"playtime,omitempty"`
	Language    string   `json:"language,omitempty"`
	Platforms   []string `json:"platforms,omitempty,omitempty"`
	Stores      []string `json:"stores,omitempty,omitempty"`
	Modes       []string `json:"modes,omitempty,omitempty"`
	Genres      []string `json:"genres,omitempty,omitempty"`
	Publishers  []string `json:"publishers,omitempty,omitempty"`
	Developers  []string `json:"developers,omitempty,omitempty"`
	Website     string   `json:"website,omitempty"`
	Tags        []string `json:"tags,omitempty,omitempty"`
	Trailer     string   `json:"trailer,omitempty"`
	AgeRating   string   `json:"age_rating,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

type PatchGame = CreateGame

type GameResponse struct {
	BaseResponse
	Data Game `json:"data"`
}

type GameListResponse struct {
	BaseResponse
	Data []Game `json:"data"`
}

// ListGames will get all games
func (c *Client) ListGames(options ...ListOptions) (GameListResponse, error) {
	var games []Game
	var res GameListResponse

	opts := c.newListOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(gameOrigin, &opts, &games)
	if err != nil {
		return res, err
	}

	res = GameListResponse{
		BaseResponse: doRes,
		Data:         games,
	}

	return res, nil
}

// FindGame will get a single item by uuid
func (c *Client) FindGame(id uuid.UUID) (GameResponse, error) {
	var game Game
	var res GameResponse

	doRes, err := c.doGetRequest(gameOrigin+"/"+id.String(), nil, &game)
	if err != nil {
		return res, err
	}

	res = GameResponse{
		BaseResponse: doRes,
		Data:         game,
	}

	return res, nil
}

func (c *Client) RandomGame() (GameResponse, error) {
	var game Game
	var res GameResponse

	doRes, err := c.doGetRequest(gameOrigin+"/random", nil, game)
	if err != nil {
		return res, err
	}

	res = GameResponse{
		BaseResponse: doRes,
		Data:         game,
	}

	return res, nil
}

func (c *Client) CreateGame(s CreateGame) (Game, error) {
	var game Game

	err := c.doPostRequest(gameOrigin, s, &game)
	if err != nil {
		return game, err
	}

	return game, nil
}

func (c *Client) PatchGame(id uuid.UUID, p PatchGame) (Game, error) {
	var game Game

	err := c.doPatchRequest(gameOrigin+"/"+id.String(), &p)
	if err != nil {
		return game, err
	}

	res, err := c.FindGame(id)
	if err != nil {
		return game, err
	}

	game = res.Data
	return game, nil
}

func (c *Client) DeleteGame(id uuid.UUID) error {
	return c.doDeleteRequest(gameOrigin + "/" + id.String())
}
