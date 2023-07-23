package service

type Category struct {
	Name string
}

func (a *app) CreateCategory(name string) error {
	return a.store.CreateCategory(name)
}

func (a *app) GetAllCategories() ([]Category, error) {
	categoryRows, err := a.store.GetAllCategories()
	if err != nil {
		return nil, err
	}

	categories := []Category{}
	for _, categoryRow := range categoryRows {
		categories = append(categories, Category{
			Name: categoryRow.Name,
		})
	}

	return categories, nil
}

func (a *app) UpdateCategory(categoryID int, name string) error {
	return a.store.UpdateCategory(categoryID, name)
}

func (a *app) DeleteCategory(categoryID int) error {
	return a.store.DeleteCategory(categoryID)
}
