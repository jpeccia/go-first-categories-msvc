package repositories

import "github.com/jpeccia/go-first-categories-msvc/internal/entities"

type InMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *InMemoryCategoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)

	return nil
}
