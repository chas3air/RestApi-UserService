package storage

type Repository[T any] interface {
	Get() ([]T, error)
	GetById(int) (T, error)
	Insert(T) error
	Update(int, T) error
	Delete(int) error
}
