package hawapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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
	if r := reflect.ValueOf(out); out != nil && r.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("out must be a pointer")
	}

	req.Header.Set("Content-Type", "application/json")

	// Token is optional
	if len(c.options.Token) != 0 {
		req.Header.Set("Authorization", "Bearer "+c.options.Token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != wantStatus {
		var resErr ErrorResponse
		if err := json.Unmarshal(body, &resErr); err != nil {
			return nil, errors.New("failed to parse error message: " + err.Error())
		}
		return nil, resErr
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return nil, err
		}
	}

	return res.Header, nil
}

func (c *Client) doGetRequest(origin string, query []QueryOptions, out any) (BaseResponse, error) {
	var res BaseResponse

	// This will fix 'buildUrl' ignoring url options if 'query' is nil
	if query == nil {
		query = []QueryOptions{}
	}

	url := c.buildUrl(origin, query)
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

func (c *Client) doPostRequest(origin string, in any, out any) error {
	if len(c.options.Token) == 0 {
		return fmt.Errorf("token is required for post request")
	}

	url := c.buildUrl(origin, nil)
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

func (c *Client) doPatchRequest(origin string, patch any) error {
	if len(c.options.Token) == 0 {
		return fmt.Errorf("token is required for put request")
	}

	var item any
	_, err := c.doGetRequest(origin, nil, &item)
	if err != nil {
		return err
	}

	res, err := json.Marshal(patch)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return err
	}

	itemBytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	url := c.buildUrl(origin, nil)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(itemBytes))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, http.StatusOK, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doDeleteRequest(origin string) error {
	if len(c.options.Token) == 0 {
		return fmt.Errorf("token is required for delete request")
	}

	url := c.buildUrl(origin, nil)
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

func (c *Client) buildUrl(origin string, query []QueryOptions) string {
	url := fmt.Sprintf("%s/%s/%s", c.options.Endpoint, c.options.Version, origin)

	// No options to append
	if query == nil {
		return url
	}

	var params []string

	// Don't set language param if it's the same as default
	if len(c.options.Language) != 0 && c.options.Language != DefaultLanguage {
		params = pushOrOverwrite(params, "language", c.options.Language)
	}

	// Don't set size param if it's the same as default
	if c.options.Size != 0 && c.options.Size != DefaultSize {
		params = pushOrOverwrite(params, "size", strconv.Itoa(c.options.Size))
	}

	opts := c.newQueryOptions()
	for _, opt := range query {
		opt(&opts)
	}

	for key, value := range opts.Filters {
		if value != "" {
			params = pushOrOverwrite(params, key, value)
		}
	}

	if opts.Pageable.Page != 0 && opts.Pageable.Page != 1 {
		params = pushOrOverwrite(params, "page", strconv.Itoa(opts.Pageable.Page))
	}

	if opts.Pageable.Size != 0 && opts.Pageable.Size != DefaultSize {
		params = pushOrOverwrite(params, "size", strconv.Itoa(opts.Pageable.Size))
	}

	if opts.Pageable.Sort != "" {
		sortParam := opts.Pageable.Sort
		if opts.Pageable.Order != "" {
			sortParam = fmt.Sprintf("%s,%s", sortParam, opts.Pageable.Order)
		}
		params = pushOrOverwrite(params, "sort", sortParam)
	}

	paramsStr := ""
	if len(params) > 0 {
		paramsStr = "?" + strings.Join(params, "&")
	}

	url += paramsStr
	return url
}

func pushOrOverwrite(params []string, key, value string) []string {
	for i, param := range params {
		if strings.HasPrefix(param, key+"=") {
			params[i] = fmt.Sprintf("%s=%s", key, value)
			return params
		}
	}
	return append(params, fmt.Sprintf("%s=%s", key, value))
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

	nextPage := handlePagination(headers.Page, true)
	headers.NextPage = nextPage

	prevPage := handlePagination(headers.Page, false)
	headers.PrevPage = prevPage

	headers.Etag = header.Get(apiHeaderEtag)
	headers.Language = header.Get(apiHeaderContentLanguage)
	return headers
}

func handlePagination(page int, increase bool) int {
	if page <= 0 {
		return -1
	}

	if increase {
		page++
	} else {
		page--
	}

	if page == 0 {
		return -1
	}

	return page
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
