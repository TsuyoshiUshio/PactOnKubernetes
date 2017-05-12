package types

import "errors"

// User is a representation of a User. Dah.
type Product struct {
	Id    int64
	Name  string `json:"name"`
	Price int64
}

var (
	// ErrNotFound represents a resource not found (404)
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized represents a Forbidden (403)
	ErrUnauthorized = errors.New("unauthorized")

	// ErrEmpty is returned when input string is empty
	ErrEmpty = errors.New("empty string")
)

type SearchRequest struct {
	Keyword string `json:"keyword"`
}

type SearchResponse struct {
	Product *Product `json:"product"`
}
