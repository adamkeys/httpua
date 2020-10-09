// Package httpua provides a HTTP client that sends a custom User-Agent.
package httpua

import "net/http"

// NewClient returns a new http.Client that is initialized to send the supplied userAgent string for the
// User-Agent header.
func NewClient(userAgent string) *http.Client {
	return WithClient(http.DefaultClient, userAgent)
}

// WithClient is similar to NewClient but returns a new http.Client copied from the supplied http.Client that is
// initialized to send the supplied userAgent string for the User-Agent header.
func WithClient(client *http.Client, userAgent string) *http.Client {
	client = &(*client) // Copy the http.Client
	client.Transport = newTransport(client.Transport, userAgent)
	return client
}

// transport implements a http.RoundTripper that sets the User-Agent for each request.
type transport struct {
	Transport http.RoundTripper
	UserAgent string
}

// newTransport returns a transport configured with the supplied values. If the roundTripper is nil,
// http.DefaultTransport will be used.
func newTransport(roundTripper http.RoundTripper, userAgent string) *transport {
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	return &transport{
		Transport: roundTripper,
		UserAgent: userAgent,
	}
}

// RoundTrip assigns the User-Agent header and then passes the work to the parent Transport.
func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", t.UserAgent)
	return t.Transport.RoundTrip(r)
}
