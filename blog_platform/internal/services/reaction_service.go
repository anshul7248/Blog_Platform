package services

import (
	"blog_project/internal/models"
	"blog_project/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
}

func NewReactionService(r *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{
		repo: r,
	}
}

func (s *ReactionService) TogglePostLike(postID, userID string) (bool, error) {
	if re, err := s.repo.Find(postID, userID); err == nil && re != nil {
		if err := s.repo.Delete(re.ID.String()); err != nil {
			return false, err
		}
		return false, nil
	}

	nr := &models.Reaction{
		ID:        uuid.New(),
		PostID:    uuid.MustParse(postID),
		UserID:    uuid.MustParse(userID),
		Type:      "like",
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(nr); err != nil {
		return false, err
	}
	return true, nil
}
