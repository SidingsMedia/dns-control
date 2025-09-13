package model

type List[T any] struct {
	Results []T `json:"results"`
}
