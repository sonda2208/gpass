package googlepasses

import "net/url"

type QueryParams map[string][]string

func (p QueryParams) Get(key string) string {
	v := p[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func (p QueryParams) Set(key, value string) {
	p[key] = []string{value}
}

func (p QueryParams) SetMulti(key string, values []string) {
	p[key] = values
}

func (p QueryParams) Encode() string {
	return url.Values(p).Encode()
}
