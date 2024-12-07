package repository

import (
	"context"
	"errors"
	"fmt"
	myContext "golizilla/adapters/http/handler/context"
	"golizilla/core/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, userCtx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, userCtx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, userCtx context.Context, id uuid.UUID) (*model.User, error)
	Update(ctx context.Context, userCtx context.Context, user *model.User) error
	// profile
	CreateNotification(ctx context.Context, userCtx context.Context, userId uuid.UUID, notification *model.Notification) error
	FindByIDWithNotifications(ctx context.Context, userCtx context.Context, userId uuid.UUID) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, userCtx context.Context, user *model.User) error {
	var db *gorm.DB
	if db = myContext.GetDB(userCtx); db == nil {
		db = r.db
	}

	return db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, userCtx context.Context, email string) (*model.User, error) {
	var db *gorm.DB
	if db = myContext.GetDB(userCtx); db == nil {
		db = r.db
	}
	var user model.User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, userCtx context.Context, id uuid.UUID) (*model.User, error) {
	var db *gorm.DB
	if db = myContext.GetDB(userCtx); db == nil {
		db = r.db
	}
	var user model.User
	err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, userCtx context.Context, user *model.User) error {
	var db *gorm.DB
	if db = myContext.GetDB(userCtx); db == nil {
		db = r.db
	}
	return db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) CreateNotification(ctx context.Context, userCtx context.Context, userId uuid.UUID, notification *model.Notification) error {
	var db *gorm.DB
	if db = myContext.GetDB(userCtx); db == nil {
		db = r.db
	}
	notification.UserID = userId
	if err := db.WithContext(ctx).Create(notification).Error; err != nil {
		return fmt.Errorf("failed to create notification: %w", err)
	}
	return nil
}

// it's need to test
func (r *UserRepository) FindByIDWithNotifications(ctx context.Context, userCtx context.Context, userId uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).
		Preload("NotificationList").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}