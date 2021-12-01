package context

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <- ctx.Done():
		return "", ctx.Err()
	case res:= <-data:
		return res, nil


	}
}


type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //TODO: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
