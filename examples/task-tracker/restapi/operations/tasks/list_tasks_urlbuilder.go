package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"

	"github.com/go-openapi/swag"
)

// ListTasksURL generates an URL for the list tasks operation
type ListTasksURL struct {
	PageSize *int32
	SinceID  *int64
	Status   []string
	Tags     []string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *ListTasksURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/tasks"

	result.Path = _path

	qs := make(url.Values)

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt32(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var sinceID string
	if o.SinceID != nil {
		sinceID = swag.FormatInt64(*o.SinceID)
	}
	if sinceID != "" {
		qs.Set("sinceId", sinceID)
	}

	var statusIR []string
	for _, statusI := range o.Status {
		statusIS := statusI
		if statusIS != "" {
			statusIR = append(statusIR, statusIS)
		}
	}

	status := swag.JoinByFormat(statusIR, "pipes")

	if len(status) > 0 {
		qsv := status[0]
		if qsv != "" {
			qs.Set("status", qsv)
		}
	}

	var tagsIR []string
	for _, tagsI := range o.Tags {
		tagsIS := tagsI
		if tagsIS != "" {
			tagsIR = append(tagsIR, tagsIS)
		}
	}

	tags := swag.JoinByFormat(tagsIR, "")

	if len(tags) > 0 {
		qsv := tags[0]
		if qsv != "" {
			qs.Set("tags", qsv)
		}
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListTasksURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListTasksURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListTasksURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListTasksURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListTasksURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListTasksURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
