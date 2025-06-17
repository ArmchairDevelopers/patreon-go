package patreon

import (
	"strings"
)

type options struct {
	fields     map[string]string
	filters    map[string]string
	include    string
	size       int
	cursor     *string
	maxResults int
}

type RequestOpt func(*options)

// WithFields specifies the resource attributes you want to be returned by API.
func WithFields(resource string, fields ...string) RequestOpt {
	return func(o *options) {
		if o.fields == nil {
			o.fields = make(map[string]string)
		}
		o.fields[resource] = strings.Join(fields, ",")
	}
}

// WithIncludes specifies the related resources you want to be returned by API.
func WithIncludes(include ...string) RequestOpt {
	return func(o *options) {
		o.include = strings.Join(include, ",")
	}
}

// WithPageSize specifies the number of items to return.
func WithPageSize(size int) RequestOpt {
	return func(o *options) {
		o.size = size
	}
}

// WithFilter allows you to filter results based on a specific field and value.
func WithFilter(field, value string) RequestOpt {
	return func(o *options) {
		if o.filters == nil {
			o.filters = make(map[string]string)
		}

		o.filters[field] = value
	}
}

// withCursor controls cursor-based pagination. Cursor will also be extracted from navigation links for convenience.
func withCursor(cursor *string) RequestOpt {
	return func(o *options) {
		o.cursor = cursor
	}
}

func getOptions(opts ...RequestOpt) options {
	cfg := options{}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}
