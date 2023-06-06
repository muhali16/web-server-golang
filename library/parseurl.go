package library

import (
	"net/url"
)

func ParseUrl(path string) map[string]interface{} {
	var u, err = url.Parse(path)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	var result = map[string]interface{}{
		"protocol": u.Scheme,
		"host":     u.Host,
		"path":     u.Path,
		"query":    u.RawQuery,
	}

	return result
}
