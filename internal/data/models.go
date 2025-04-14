package data

import (
	"database/sql"
)

type Models struct {
	Comments CommentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Comments: CommentModel{
			DB: db,
		},
	}
}
