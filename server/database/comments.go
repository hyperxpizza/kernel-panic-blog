package database

import (
	"time"

	"github.com/hyperxpizza/kernel-panic-blog/server/models"
)

func GetCommentsByPostID(postID int) ([]*models.Comment, error) {
	var comments []*models.Comment
	rows, err := db.Query(`SELECT * FROM comments WHERE postID=$1`, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment

		err = rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.IsAdmin,
			&comment.OPEmail,
			&comment.OPName,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func InsertComment(postID int, isAdmin bool, content, opEmail, opName string) error {
	stmt, err := db.Prepare(`INSERT INTO comments VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(postID, content, time.Now(), time.Now(), isAdmin, opEmail, opName)
	if err != nil {
		return err
	}

	return nil
}
