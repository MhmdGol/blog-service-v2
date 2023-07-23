package catlisthandler

import "blog-service/internal/repository"

func CategoriesToList(catListModel []*repository.Category) (catListString []string) {
	for _, cat := range catListModel {
		catListString = append(catListString, cat.Name)
	}
	return catListString
}
