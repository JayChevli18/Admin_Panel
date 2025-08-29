package user

import (
	context "context"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	Create(ctx context.Context, req *CreateUserRequest) (*User, error)
	Get(ctx context.Context, userId int64) (*User, error)
	Update(ctx context.Context, userId int64, req UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, userId int64) error
	List(ctx context.Context, page, pageSize int) ([]User, int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	user := &User{
		UserID:    req.UserID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Role:      Role(req.Role),
		IsActive:  *req.IsActive,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Update(ctx context.Context, userId int64, req UpdateUserRequest) (*User, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Role != nil {
		user.Role = Role(*req.Role)
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.repo.UpdateByUserId(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Get(ctx context.Context, userId int64) (*User, error) {
	return s.repo.GetByUserId(ctx, userId)
}

func (s *service) Delete(ctx context.Context, userId int64) error {
	return s.repo.DeleteByUserId(ctx, userId)
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]User, int64, error) {
	return s.repo.List(ctx, page, pageSize)
}
