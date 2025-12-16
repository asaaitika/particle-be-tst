package repository

import (
	"article-api/internal/models"
	"database/sql"
	"fmt"
)

type PostRepository interface {
	Create(post *models.CreatePostRequest) (*models.Post, error)
	GetAll(limit, offset int) ([]models.Post, error)
	GetByID(id int) (*models.Post, error)
	Update(id int, post *models.UpdatePostRequest) (*models.Post, error)
	Delete(id int) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(req *models.CreatePostRequest) (*models.Post, error) {
	query := `
		INSERT INTO posts (title, content, category, status, created_date, updated_date)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, title, content, category, status, created_date, updated_date
	`

	post := &models.Post{}
	err := r.db.QueryRow(query, req.Title, req.Content, req.Category, req.Status).
		Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Status, &post.CreatedDate, &post.UpdatedDate)

	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return post, nil
}

func (r *postRepository) GetAll(limit, offset int) ([]models.Post, error) {
	query := `
		SELECT id, title, content, category, status, created_date, updated_date
		FROM posts
		ORDER BY created_date DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Status, &post.CreatedDate, &post.UpdatedDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}

	if posts == nil {
		posts = []models.Post{}
	}

	return posts, nil
}

func (r *postRepository) GetByID(id int) (*models.Post, error) {
	query := `
		SELECT id, title, content, category, status, created_date, updated_date
		FROM posts
		WHERE id = $1
	`

	post := &models.Post{}
	err := r.db.QueryRow(query, id).
		Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Status, &post.CreatedDate, &post.UpdatedDate)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return post, nil
}

func (r *postRepository) Update(id int, req *models.UpdatePostRequest) (*models.Post, error) {
	query := `
		UPDATE posts
		SET title = $1, content = $2, category = $3, status = $4, updated_date = NOW()
		WHERE id = $5
		RETURNING id, title, content, category, status, created_date, updated_date
	`

	post := &models.Post{}
	err := r.db.QueryRow(query, req.Title, req.Content, req.Category, req.Status, id).
		Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Status, &post.CreatedDate, &post.UpdatedDate)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return post, nil
}

func (r *postRepository) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}
