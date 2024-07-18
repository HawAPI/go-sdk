package hawapi

import (
	"net/http"
	"time"

	"github.com/HawAPI/go-sdk/pkg/cache"
)

const (
	DefaultEndpoint         = "https://hawapi.theproject.id/api"
	DefaultVersion          = "v1"
	DefaultLanguage         = "en-US"
	DefaultSize             = 10
	DefaultTimeout          = 10
	DefaultUseInMemoryCache = true
)

// DefaultOptions for Go HawAPI SDK
var DefaultOptions = Options{
	Endpoint:         DefaultEndpoint,
	Version:          DefaultVersion,
	Language:         DefaultLanguage,
	Size:             DefaultSize,
	Timeout:          DefaultTimeout,
	UseInMemoryCache: DefaultUseInMemoryCache,
}

type Options struct {
	// The endpoint of the HawAPI instance
	//
	// Default value: DefaultEndpoint
	Endpoint string

	// The version of the API
	Version string

	// The language of items for all requests
	//
	// Note: This value can be overwritten later
	Language string

	// The size of items for all requests
	//
	// Note: This value can be overwritten later
	Size int

	// The timeout of a response in milliseconds
	Timeout int

	// The HawAPI token (JWT)
	//
	// By default, all requests are made with 'ANONYMOUS' tier
	Token string

	// Define if the package should save (in-memory) all request results
	UseInMemoryCache bool
}

// Client is the [HawAPI] golang client.
//
//   - [GitHub]
//   - [Examples]
//
// [HawAPI]: https://github.com/HawAPI/HawAPI
// [GitHub]: https://github.com/HawAPI/go-sdk/
// [Examples]: https://github.com/HawAPI/go-sdk/examples/
type Client struct {
	options Options
	client  *http.Client
	cache   cache.Cache
}

// NewClient creates a new HawAPI client using the default options.
func NewClient() Client {
	return NewClientWithOpts(DefaultOptions)
}

// NewClientWithOpts creates a new HawAPI client using custom options.
func NewClientWithOpts(options Options) Client {
	c := Client{}
	c.WithOpts(options)

	c.client = &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}

	c.cache = cache.NewMemoryCache()
	return c
}

// WithOpts will set or override current client options
func (c *Client) WithOpts(options Options) {
	if len(options.Endpoint) != 0 {
		c.options.Endpoint = options.Endpoint
	}

	if len(options.Version) != 0 {
		c.options.Version = options.Version
	}

	if len(options.Language) != 0 {
		c.options.Language = options.Language
	}

	if options.Size != 0 {
		c.options.Size = options.Size
	}

	if options.Timeout != 0 {
		c.options.Timeout = options.Timeout
	}

	c.options.UseInMemoryCache = options.UseInMemoryCache
}

// ClearCache deletes all values from the cache and returns the count of deleted items
func (c *Client) ClearCache() int {
	return c.cache.Clear()
}

// CacheSize returns the count cache items
func (c *Client) CacheSize() int {
	return c.cache.Size()
}
