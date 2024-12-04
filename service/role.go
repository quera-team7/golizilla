package service

import (
	"context"
	"golizilla/domain/model"
	"golizilla/domain/repository"
	"time"

	"github.com/google/uuid"
)

type IRoleService interface {
	CreateRole(ctx context.Context, name, description string) (*model.Role, error)
	GetRoleById(ctx context.Context, id uuid.UUID) (*model.Role, error)
	GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*model.Role, error)
	HasPrivileges(ctx context.Context, id uuid.UUID, privileges ...string) (bool, error)
}

type roleService struct {
	roleRepo          repository.IRoleRepository
	userRepo          repository.IUserRepository
	rolePrivilegeRepo repository.IRolePrivilegeRepository
}

func NewRoleService(roleRepo repository.IRoleRepository, userRepo repository.IUserRepository, rolePrivilegeRepo repository.IRolePrivilegeRepository) IRoleService {
	return &roleService{
		roleRepo:          roleRepo,
		userRepo:          userRepo,
		rolePrivilegeRepo: rolePrivilegeRepo,
	}
}

func (s *roleService) CreateRole(ctx context.Context, name, description string) (*model.Role, error) {
	role := &model.Role{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
	if err := s.roleRepo.Add(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleService) GetRoleById(ctx context.Context, id uuid.UUID) (*model.Role, error) {
	role, err := s.roleRepo.GetById(ctx, id)
	if err != nil {
		//log
	}
	return role, err
}

func (s *roleService) GetRoleByUserId(ctx context.Context, userId uuid.UUID) (*model.Role, error) {
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		//log
		return nil, err
	}
	role, err := s.roleRepo.GetById(ctx, user.RoleId)
	if err != nil {
		//log
		return nil, err
	}
	return role, err
}

func (s *roleService) HasPrivileges(ctx context.Context, id uuid.UUID, privileges ...string) (bool, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		//log
		return false, err
	}
	rolePrivileges, err := s.rolePrivilegeRepo.GetRolePrivilegesByPrivileges(ctx, user.RoleId, privileges...)
	if err != nil {
		//log
		return false, err
	}

	if len(rolePrivileges) == len(privileges) {
		return true, nil
	} else {
		return false, nil
	}
}