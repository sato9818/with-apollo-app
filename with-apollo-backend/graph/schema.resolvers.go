package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "test-with-docker/util"
	"log"
	"sort"
	// "strconv"
	"time"
	"fmt"

	"test-with-docker/entity"
	"test-with-docker/graph/generated"
	"test-with-docker/graph/model"
)

// var posts []*model.Post = make([]*model.Post, 0)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, url string) (*model.Post, error) {
	post := entity.Post{
		Title:     title,
		URL:       url,
		Votes:     0,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	result := r.DB.Create(&post); 
	err := result.Error;
	if err != nil {
		return nil, err
	}
	new_post := model.Post{
		ID:        fmt.Sprintf("%d", post.ID),
		Title:     post.Title,
		URL:       post.URL,
		Votes:     0,
		CreatedAt: post.CreatedAt,
	}
	log.Printf("%#v\n", new_post)
	return &new_post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, votes *int) (*model.Post, error) {
	if votes == nil {
		return nil, nil
	}
	var post model.Post; 
	result := r.DB.First(&post, id);
	err := result.Error;
	if err != nil {
		return nil, err
	}

	post.Votes = *votes + 1
	r.DB.Save(&post)
	return &post, nil
}

func (r *queryResolver) AllPosts(ctx context.Context, orderBy *model.OrderBy, first int, skip int) ([]*model.Post, error) {
	
	var posts []*model.Post;
	result := r.DB.Find(&posts);
	err := result.Error;
	if err != nil {
		return nil, err
	}
	if skip > len(posts) {
		skip = len(posts)
	}
	if (skip + first) > len(posts) {
		first = len(posts) - skip
	}
	sortedPosts := make([]*model.Post, len(posts))
	copy(sortedPosts, posts)
	if orderBy != nil && *orderBy == "createdAt_DESC" {
		sort.SliceStable(sortedPosts, func(i, j int) bool {
			return sortedPosts[i].CreatedAt > sortedPosts[j].CreatedAt
		})
	}
	slicePosts := sortedPosts[skip : skip+first]
	return slicePosts, nil
}

func (r *queryResolver) AllPostsMeta(ctx context.Context) (*model.PostsMeta, error) {
	var posts []*model.Post;
	result := r.DB.Find(&posts);
	err := result.Error;
	if err != nil {
		return nil, err
	}
	postsMeta := model.PostsMeta{Count: len(posts)}
	return &postsMeta, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
