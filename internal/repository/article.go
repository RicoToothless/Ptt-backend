package repository

import (
	"context"

	"github.com/Ptt-official-app/go-bbs"
)

// PopularArticleRecord is an ArticleRecord which has boardID information.
type PopularArticleRecord interface {
	// Note: go-bbs has not implemented this yet
	// TODO: use bbs.PopularArticleRecord or something when it is ready
	bbs.ArticleRecord
	BoardID() string
}

func (repo *repository) GetPopularArticles(ctx context.Context) ([]PopularArticleRecord, error) {
	// Note: go-bbs has not implemented this yet
	// TODO: delegate to repo.db when it is ready
	return []PopularArticleRecord{}, nil
}

func (repo *repository) AppendComment(ctx context.Context, userID, boardID, filename, appendType, text string) (map[string]interface{}, error) {
	return nil, nil
}

func (repo *repository) ForwardArticleToBoard(ctx context.Context, userID, boardID, filename, boardName string) (map[string]interface{}, error) {
	result := map[string]interface{}{
		"user_id":      "No value",
		"article_id":   "No value",
		"forward_time": "No value",
		"title":        "No value",
		"board_id":     "No value",
	}
	return result, nil
}

func (repo *repository) ForwardArticleToEmail(ctx context.Context, userID string, boardID string, filename string, email string) error {
	return nil
}
