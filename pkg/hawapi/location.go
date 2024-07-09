package hawapi

import (
	"github.com/google/uuid"
)

const locationOrigin = "locations"

type Location struct {
	Uuid        uuid.UUID `json:"uuid"`
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Images      []string  `json:"images,omitempty"`
	Sources     []string  `json:"sources,omitempty"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type CreateLocation struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Language    string   `json:"language,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	Images      []string `json:"images,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

type PatchLocation = CreateLocation

type LocationResponse struct {
	BaseResponse
	Data Location `json:"data"`
}

type LocationListResponse struct {
	BaseResponse
	Data []Location `json:"data"`
}

// ListLocations will get all locations
func (c *Client) ListLocations(options ...QueryOptions) (LocationListResponse, error) {
	var locations []Location
	var res LocationListResponse

	opts := c.newQueryOptions()
	for _, opt := range options {
		opt(&opts)
	}

	doRes, err := c.doGetRequest(locationOrigin, &opts, &locations)
	if err != nil {
		return res, err
	}

	res = LocationListResponse{
		BaseResponse: doRes,
		Data:         locations,
	}

	return res, nil
}

// FindLocation will get a single item by uuid
func (c *Client) FindLocation(id uuid.UUID) (LocationResponse, error) {
	var location Location
	var res LocationResponse

	opts := c.newQueryOptions()
	doRes, err := c.doGetRequest(locationOrigin+"/"+id.String(), &opts, &location)
	if err != nil {
		return res, err
	}

	res = LocationResponse{
		BaseResponse: doRes,
		Data:         location,
	}

	return res, nil
}

func (c *Client) RandomLocation() (LocationResponse, error) {
	var location Location
	var res LocationResponse

	opts := c.newQueryOptions()
	doRes, err := c.doGetRequest(locationOrigin+"/random", &opts, &location)
	if err != nil {
		return res, err
	}

	res = LocationResponse{
		BaseResponse: doRes,
		Data:         location,
	}

	return res, nil
}

func (c *Client) CreateLocation(s CreateLocation) (Location, error) {
	var location Location

	err := c.doPostRequest(locationOrigin, s, &location)
	if err != nil {
		return location, err
	}

	return location, nil
}

func (c *Client) PatchLocation(id uuid.UUID, p PatchLocation) (Location, error) {
	var location Location

	err := c.doPatchRequest(locationOrigin+"/"+id.String(), &p)
	if err != nil {
		return location, err
	}

	res, err := c.FindLocation(id)
	if err != nil {
		return location, err
	}

	location = res.Data
	return location, nil
}

func (c *Client) DeleteLocation(id uuid.UUID) error {
	return c.doDeleteRequest(locationOrigin + "/" + id.String())
}
