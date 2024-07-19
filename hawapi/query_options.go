package hawapi

type Filters map[string]string

type queryOptions struct {
	Pageable
	Filters
}

type QueryOptions func(*queryOptions)

func NewQueryOptions(pageable Pageable, filters Filters) QueryOptions {
	return func(o *queryOptions) {
		o.Pageable = pageable
		o.Filters = filters
	}
}

// newQueryOptions wil create a new queryOptions with default and pre-defined values.
//
// Values like 'page size' and 'language' are configured on Client initialization
func (c *Client) newQueryOptions() queryOptions {
	opts := queryOptions{
		Pageable: Pageable{
			Page:  1,
			Size:  DefaultSize,
			Sort:  "",
			Order: "ASC",
		},
		Filters: make(Filters),
	}

	return opts
}

func WithFilters(filters Filters) QueryOptions {
	return func(o *queryOptions) {
		o.Filters = filters
	}
}

func WithFilter(key string, value string) QueryOptions {
	return func(o *queryOptions) {
		o.Filters[key] = value
	}
}

func WithLanguage(language string) QueryOptions {
	return WithFilter("language", language)
}

func WithPageable(pageable Pageable) QueryOptions {
	return func(o *queryOptions) {
		o.Pageable = pageable
	}
}

func WithPage(page int) QueryOptions {
	return func(o *queryOptions) {
		o.Page = page
	}
}

func WithSize(size int) QueryOptions {
	return func(o *queryOptions) {
		o.Size = size
	}
}

func WithSort(sort string) QueryOptions {
	return func(o *queryOptions) {
		o.Sort = sort
	}
}

func WithOrder(order string) QueryOptions {
	return func(o *queryOptions) {
		o.Order = order
	}
}
