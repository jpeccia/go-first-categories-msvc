package usecases

import (
	"log"

	"github.com/jpeccia/go-first-categories-msvc/internal/entities"
)

type createCategoryUseCase struct {
	// db
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)

	return nil
}
