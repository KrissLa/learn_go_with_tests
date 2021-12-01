package context

import (
	"errors"
	"net/http"
)

type SpyResponseWritter struct {
	written bool
}

func (s *SpyResponseWritter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWritter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWritter) WriteHeader(statusCode int) {
	s.written = true
}
