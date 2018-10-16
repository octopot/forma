package domain

import (
	"net/http"
	"strings"
)

const (
	cookieHeader    = "Cookie"
	optionsHeader   = "X-Forma-Options"
	refererHeader   = "Referer"
	userAgentHeader = "User-Agent"
)

// FromCookies returns converted value from request's cookies.
func FromCookies(cookies []*http.Cookie) map[string]string {
	converted := make(map[string]string)
	for _, cookie := range cookies {
		if cookie.HttpOnly && cookie.Secure {
			converted[cookie.Name] = cookie.Value
		}
	}
	return converted
}

// FromHeaders returns converted value from request's headers.
func FromHeaders(headers http.Header) map[string][]string {
	converted := make(map[string][]string)
	for key, values := range headers {
		if !strings.EqualFold(cookieHeader, key) {
			converted[key] = values
		}
	}
	return converted
}

// FromRequest returns converted value from a request.
func FromRequest(req *http.Request) map[string][]string {
	return req.URL.Query()
}

// Header TODO issue#173
func (context InputContext) Header(key string) string {
	return http.Header(context.Headers).Get(key)
}

// Option TODO issue#173
func (context InputContext) Option() Option {
	split := func(str string) []string {
		ss := strings.Split(str, ";")
		for i, s := range ss {
			ss[i] = strings.ToLower(strings.TrimSpace(s))
		}
		return ss
	}
	is := func(where []string, what string) bool {
		for _, opt := range where {
			if opt == what {
				return true
			}
		}
		return false
	}
	options := split(context.Header(optionsHeader))
	return Option{
		Anonymously: is(options, "anonym"),
		Debug:       is(options, "debug"),
		NoLog:       is(options, "nolog"),
	}
}

// Referer TODO issue#173
func (context InputContext) Referer() string {
	return context.Header(refererHeader)
}

// UserAgent TODO issue#173
func (context InputContext) UserAgent() string {
	return context.Header(userAgentHeader)
}
