package service

import (
	"context"
	"golizilla/domain/model"
	"golizilla/domain/repository"

	"github.com/google/uuid"
)

type IAnswerService interface {
	Create(ctx context.Context, answer *model.Answer) (uuid.UUID, error)
	Update(ctx context.Context, answer *model.Answer) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Answer, error)
}

type AnswerService struct {
	answerRepo repository.IAnswerRepository
}

func NewAnswerService(repo repository.IAnswerRepository) IAnswerService {
	return &AnswerService{
		answerRepo: repo,
	}
}

func (s *AnswerService) Create(ctx context.Context, answer *model.Answer) (uuid.UUID, error) {
	return s.answerRepo.Create(ctx, answer)
}

func (s *AnswerService) Update(ctx context.Context, answer *model.Answer) error {
	return s.answerRepo.Update(ctx, answer)
}

func (s *AnswerService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.answerRepo.Delete(ctx, id)
}

func (s *AnswerService) GetByID(ctx context.Context, id uuid.UUID) (*model.Answer, error) {
	return s.answerRepo.GetByID(ctx, id)
}