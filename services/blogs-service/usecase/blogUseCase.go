package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPostUseCase interface {
	GetAllBlogPosts(ctx context.Context, limit, skip int, sortField, sortOrder string) ([]domain.BlogPost, error)
	GetBlogPostByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogPost, error)
	GetBlogPostByTitle(ctx context.Context, title string) (*domain.BlogPost, error)
	CreateBlogPost(ctx context.Context, post *domain.BlogPost) error
	UpdateBlogPost(ctx context.Context, post *domain.BlogPost) error
	DeleteBlogPost(ctx context.Context, id primitive.ObjectID) error
}

type blogPostUseCase struct {
	blogPostRepository repository.BlogPostRepository
}

func NewBlogPostUseCase(repository repository.BlogPostRepository) *blogPostUseCase {
	return &blogPostUseCase{
		blogPostRepository: repository,
	}
}

func (u *blogPostUseCase) GetAllBlogPosts(ctx context.Context, limit, skip int, sortField, sortOrder string) ([]domain.BlogPost, error) {
	return u.blogPostRepository.GetAllBlogPosts(ctx, limit, skip, sortField, sortOrder)
}

func (u *blogPostUseCase) GetBlogPostByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogPost, error) {
	post, err := u.blogPostRepository.GetBlogPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("blog post not found")
	}
	return post, nil
}

func (u *blogPostUseCase) GetBlogPostByTitle(ctx context.Context, title string) (*domain.BlogPost, error) {
	post, err := u.blogPostRepository.GetBlogPostByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("blog post not found")
	}
	return post, nil
}

func (u *blogPostUseCase) CreateBlogPost(ctx context.Context, post *domain.BlogPost) error {
	existingPost, err := u.blogPostRepository.GetBlogPostByTitle(ctx, post.Title)
	if err != nil {
		return err
	}
	if existingPost != nil {
		return errors.New("blog post with this title already exists")
	}

	return u.blogPostRepository.CreateBlogPost(ctx, post)
}

func (u *blogPostUseCase) UpdateBlogPost(ctx context.Context, post *domain.BlogPost) error {
	return u.blogPostRepository.UpdateBlogPost(ctx, post)
}

func (u *blogPostUseCase) DeleteBlogPost(ctx context.Context, id primitive.ObjectID) error {
	return u.blogPostRepository.DeleteBlogPost(ctx, id)
}
