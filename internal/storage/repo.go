package storage

import "context"

type Repository[T any] interface {
	Get(context.Context) ([]T, error)
	GetById(context.Context, int) (T, error)
	Insert(context.Context, T) error
	Update(context.Context, int, T) error
	Delete(context.Context, int) error
}
