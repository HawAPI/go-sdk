package hawapi

import "net/http"

type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Url         string `json:"url"`
	Docs        string `json:"docs"`
	Github      string `json:"github"`
	License     string `json:"license"`
	GithubHome  string `json:"github_home"`
	ApiUrl      string `json:"api_url"`
	ApiVersion  string `json:"api_version"`
	ApiPath     string `json:"api_path"`
	ApiBaseUrl  string `json:"api_base_url"`
	LicenseUrl  string `json:"license_url"`
}

func (c *Client) Info() (Info, error) {
	var info Info

	req, err := http.NewRequest(http.MethodGet, c.options.Endpoint, nil)
	if err != nil {
		return info, err
	}

	_, err = c.doRequest(req, http.StatusOK, &info)
	if err != nil {
		return info, err
	}

	return info, nil
}
