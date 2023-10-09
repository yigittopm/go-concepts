package context2

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, store.Fetch())
	}
}
