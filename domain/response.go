package domain

type MultipleResponse[T any] struct {
	Data []T `json:"data"`
}
