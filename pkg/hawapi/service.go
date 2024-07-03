package hawapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) get(origin string, id uuid.UUID, out any) error {
	url := buildUrl(origin, id, c.options)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	setHeaders(req, c.options.Token)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintln(res.StatusCode))
	}

	if err = json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}

func setHeaders(req *http.Request, token string) {
	req.Header.Add("User-Agent", "GoSDK")

	if len(token) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

func buildUrl(origin string, id uuid.UUID, options Options) string {
	var i string
	if id != uuid.Nil {
		i = id.String()
	}

	return fmt.Sprintf("%s/%s/%s/%s", options.Endpoint, options.Version, origin, i)
}
