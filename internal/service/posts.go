package service

import (
	"context"

	catlisthandler "blog-service/internal/service/catListHandler"
)

type Post struct {
	Title      string
	Text       string
	Categories []string
}

func (a *app) CreatePost(ctx context.Context, title, text string, categories []string) error {
	return a.store.CreatePost(ctx, title, text, categories)
}

func (a *app) GetAllPosts(ctx context.Context) ([]Post, error) {
	postRows, err := a.store.GetAllPosts(ctx)
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

func (a *app) GetPagePosts(ctx context.Context, pageNumber, pageSize int) ([]Post, error) {
	postRows, err := a.store.GetPagePosts(ctx, pageNumber, pageSize)
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

func (a *app) UpdatePost(ctx context.Context, postID int, title, text string, categories []string) error {
	return a.store.UpdatePost(ctx, postID, title, text, categories)
}

func (a *app) DeletePost(ctx context.Context, postID int) error {
	return a.store.DeletePost(ctx, postID)
}
