package hawapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	// ApiHeaderRateLimitRemaining is the API rate limit remaining
	apiHeaderRateLimitRemaining = "X-Rate-Limit-Remaining"

	// ApiHeaderPageIndex is the API page index header
	apiHeaderPageIndex = "X-Pagination-Page-Index"

	// ApiHeaderPageSize is the API page size header
	apiHeaderPageSize = "X-Pagination-Page-Size"

	// ApiHeaderPageTotal is the API page total header
	apiHeaderPageTotal = "X-Pagination-Page-Total"

	// ApiHeaderItemTotal is the API item total header
	apiHeaderItemTotal = "X-Pagination-Item-Total"

	// ApiHeaderContentLanguage is the API language header
	apiHeaderContentLanguage = "Content-Language"

	// ApiHeaderContentLength is the API content length
	apiHeaderContentLength = "Content-Length"

	// ApiHeaderEtag is the API content etag
	apiHeaderEtag = "ETag"
)

func (c *Client) doRequest(req *http.Request, wantStatus int, out any) (http.Header, error) {
	req.Header.Set("Content-Type", "application/json")

	// Token is optional
	if len(c.options.Token) != 0 {
		req.Header.Set("Authorization", "Bearer "+c.options.Token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != wantStatus {
		fmt.Println(string(body))
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return nil, err
		}
	}

	return res.Header, nil
}

func (c *Client) doGetRequest(url string, out any) (BaseResponse, error) {
	var res BaseResponse

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}

	httpHeader, err := c.doRequest(req, http.StatusOK, out)
	if err != nil {
		return res, err
	}

	headers := extractHeaders(httpHeader)
	res = BaseResponse{
		HeaderResponse: headers,
		Cached:         false,
	}

	return res, nil
}

func (c *Client) doPostRequest(url string, in any, out any) error {
	if len(c.options.Token) == 0 {
		return fmt.Errorf("token is required for post request")
	}

	body, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, http.StatusCreated, out)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doDeleteRequest(url string) error {
	if len(c.options.Token) == 0 {
		return fmt.Errorf("token is required for delete request")
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, http.StatusNoContent, nil)
	if err != nil {
		return err
	}

	return nil
}

func extractHeaders(header http.Header) HeaderResponse {
	var headers HeaderResponse

	rateLimitRemaining := header.Get(apiHeaderRateLimitRemaining)
	headers.Quota.Remaining = parseInt(rateLimitRemaining)

	pageStr := header.Get(apiHeaderPageIndex)
	headers.Page = parseInt(pageStr)

	pageSizeStr := header.Get(apiHeaderPageSize)
	headers.PageSize = parseInt(pageSizeStr)

	pageTotalStr := header.Get(apiHeaderPageTotal)
	headers.PageTotal = parseInt(pageTotalStr)

	itemStr := header.Get(apiHeaderItemTotal)
	headers.ItemSize = parseInt(itemStr)

	lengthStr := header.Get(apiHeaderContentLength)
	headers.Length = parseInt(lengthStr)

	headers.Etag = header.Get(apiHeaderEtag)
	headers.Language = header.Get(apiHeaderContentLanguage)
	return headers
}

func parseInt(s string) int {
	if len(s) == 0 {
		return -1
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return i
}
