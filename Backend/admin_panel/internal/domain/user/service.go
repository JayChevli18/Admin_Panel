package user

import (
	context "context"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Service interface {
	Create(ctx context.Context, req *CreateUserRequest) (*User, error)
	Get(ctx context.Context, id primitive.ObjectID) (*User, error)
	Update(ctx context.Context, id primitive.ObjectID, req *UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
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

func (s *service) Get(ctx context.Context, id primitive.ObjectID) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id primitive.ObjectID, req *UpdateUserRequest) (*User, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.repo.GetByID(ctx, id)
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

	updatedUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *service) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]User, int64, error) {
	return s.repo.List(ctx, page, pageSize)
}
