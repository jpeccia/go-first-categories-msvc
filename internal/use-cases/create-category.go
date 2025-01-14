package usecases

import (
	"log"

	"github.com/jpeccia/go-first-categories-msvc/internal/entities"
	"github.com/jpeccia/go-first-categories-msvc/internal/repositories"
)

type createCategoryUseCase struct {
	// db
	repository repositories.InMemoryCategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.InMemoryCategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: persist-entity-to-db
	log.Println(category)
	u.repository.Save(category)

	return nil
}
