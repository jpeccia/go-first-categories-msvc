package repositories

type iCategoryRepository interface {
	Save(name string) error
}