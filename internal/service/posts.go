package service

import (
	catlisthandler "blog-service/internal/service/catListHandler"
)

type Post struct {
	Title      string
	Text       string
	Categories []string
}

func (a *app) CreatePost(title, text string, categories []string) error {
	return a.store.CreatePost(title, text, categories)
}

func (a *app) GetAllPosts() ([]Post, error) {
	postRows, err := a.store.GetAllPosts()
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Title:      postRow.Title,
			Text:       postRow.Text,
			Categories: catlisthandler.CategoriesToList(postRow.Categories),
		})
	}

	return posts, nil
}

func (a *app) GetPagePosts(pageNumber, pageSize int) ([]Post, error) {
	postRows, err := a.store.GetPagePosts(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Title:      postRow.Title,
			Text:       postRow.Text,
			Categories: catlisthandler.CategoriesToList(postRow.Categories),
		})
	}

	return posts, nil
}

func (a *app) UpdatePost(postID int, title, text string, categories []string) error {
	return a.store.UpdatePost(postID, title, text, categories)
}

func (a *app) DeletePost(postID int) error {
	return a.store.DeletePost(postID)
}
