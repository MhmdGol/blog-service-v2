package service

import "context"

type Category struct {
	Name string
}

func (a *app) CreateCategory(ctx context.Context, name string) error {
	return a.store.CreateCategory(ctx, name)
}

func (a *app) GetAllCategories(ctx context.Context) ([]Category, error) {
	categoryRows, err := a.store.GetAllCategories(ctx)
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

func (a *app) UpdateCategory(ctx context.Context, categoryID int, name string) error {
	return a.store.UpdateCategory(ctx, categoryID, name)
}

func (a *app) DeleteCategory(ctx context.Context, categoryID int) error {
	return a.store.DeleteCategory(ctx, categoryID)
}
