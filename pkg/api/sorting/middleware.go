package sorting

import (
	"context"
	"net/http"
	"strings"
)

const (
	SortASC           = "asc"
	SortDESC          = "desc"
	OptionsContextKey = "sort_options"
)

func Middleware(h http.HandlerFunc, defaultSortField, defaultSortOrder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sortBy := r.URL.Query().Get("sort_by")
		sortOrder := r.URL.Query().Get("sort_order")

		if sortBy == "" {
			sortBy = defaultSortField
		}
		if sortOrder == "" {
			sortOrder = defaultSortOrder
		}

		lowerSortOrder := strings.ToLower(sortOrder)
		if lowerSortOrder != SortASC && lowerSortOrder != SortDESC {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad sort order"))
			return
		}

		options := Options{
			Field: sortBy,
			Order: sortOrder,
		}
		ctx := context.WithValue(r.Context(), OptionsContextKey, options)
		r = r.WithContext(ctx)

		h(w, r)
	}
}

type Options struct {
	Field, Order string
}
