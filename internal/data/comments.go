package data

import (
	"database/sql"
	"errors"
	"time"
)

type Comment struct {
	ID         int64     `json:"id"`
	TimePosted time.Time `json:"-"`
	SenderName string    `json:"senderName"`
	Message    string    `json:"message"`
	Email      string    `json:"email,omitempty"`
}

type CommentModel struct {
	DB *sql.DB
}

func (c CommentModel) Insert(comment *Comment) error {
	query := `
	INSERT INTO comments (sender_name, message, email)
	VALUES ($1, $2, $3)
	RETURNING id, time_posted`

	args := []interface{}{comment.SenderName, comment.Message, comment.Email}
	return c.DB.QueryRow(query, args...).Scan(&comment.ID, &comment.TimePosted)
}

func (c CommentModel) Get(id int64) (*Comment, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
	SELECT id, time_posted, sender_name, message, email
	FROM comments
	WHERE id = $1`

	var comment Comment

	err := c.DB.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.TimePosted,
		&comment.SenderName,
		&comment.Message,
		&comment.Email,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}
	return &comment, nil
}

func (c CommentModel) Delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
	DELETE FROM comments
	WHERE id = $1
	`
	results, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (c CommentModel) GetAll() ([]*Comment, error) {
	query := `
	SELECT *
	FROM comments
	ORDER BY id`

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	comments := []*Comment{}

	for rows.Next() {
		var comment Comment

		err := rows.Scan(
			&comment.ID,
			&comment.TimePosted,
			&comment.SenderName,
			&comment.Message,
			&comment.Email,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
