package hawapi

type Filters map[string]string

type listOptions struct {
	Pageable
	Filters
}

type ListOptions func(*listOptions)

func NewListOptions(pageable Pageable, filters Filters) ListOptions {
	return func(o *listOptions) {
		o.Pageable = pageable
		o.Filters = filters
	}
}

// newListOptions wil create a new listOptions with default and pre-defined values.
//
// Values like 'page size' and 'language' are configured on Client initialization
func (c *Client) newListOptions() listOptions {
	opts := listOptions{
		Pageable: Pageable{
			Page:  1,
			Size:  c.options.Size,
			Sort:  "",
			Order: "ASC",
		},
		Filters: make(Filters),
	}

	// Don't set language param if it's the same as default
	if len(c.options.Language) != 0 && c.options.Language != DefaultLanguage {
		opts.Filters["language"] = c.options.Language
	}

	return opts
}

func WithFilters(filters Filters) ListOptions {
	return func(o *listOptions) {
		o.Filters = filters
	}
}

func WithFilter(key string, value string) ListOptions {
	return func(o *listOptions) {
		o.Filters[key] = value
	}
}

func WithLanguage(language string) ListOptions {
	return WithFilter("language", language)
}

func WithPageable(pageable Pageable) ListOptions {
	return func(o *listOptions) {
		o.Pageable = pageable
	}
}

func WithPage(page int) ListOptions {
	return func(o *listOptions) {
		o.Page = page
	}
}

func WithSize(size int) ListOptions {
	return func(o *listOptions) {
		o.Size = size
	}
}

func WithSort(sort string) ListOptions {
	return func(o *listOptions) {
		o.Sort = sort
	}
}

func WithOrder(order string) ListOptions {
	return func(o *listOptions) {
		o.Order = order
	}
}
