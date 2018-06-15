package transport

import (
	"net/http"
)

type BasicAuthTransport struct {
	Username      string
	Password      string
	WrapTransport http.RoundTripper
}

func NewBasicAuthTransport(username, password string, transport http.RoundTripper) *BasicAuthTransport {
	return &BasicAuthTransport{
		Username:      username,
		Password:      password,
		WrapTransport: transport,
	}
}

func NewDefaultBasicAuthTransport(username, password string) *BasicAuthTransport {
	return NewBasicAuthTransport(username, password, http.DefaultTransport)
}

func (t BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return t.WrapTransport.RoundTrip(req)
}
