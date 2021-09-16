package providers

import (
	"net/http"
	"regexp"
)

type SurrogateInterface interface {
	getHeaderSeparator() string
	getOrderedSurrogateKeyHeadersCandidate() []string
	getOrderedSurrogateControlHeadersCandidate() []string
	getSurrogateControl(http.Header) string
	getSurrogateKey(http.Header) string
	Purge(http.Header) []string
	purgeTag(string) []string
	Store(*http.Header, string) error
	storeTag(string, string, *regexp.Regexp)
	ParseHeaders(string) []string
	candidateStore(string) bool
}