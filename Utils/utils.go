package Utils

import (
	"go.uber.org/zap"
	"net/url"
	"strings"
)

func PreintForem(values url.Values) {
	for key, value := range values {
		zap.String(key, strings.Join(value, ","))
	}
}
