package services

import (
	"context"

	"gef/pkg/model" // 假设 model 包已正确导入
)

// UserService 用户服务接口
type UserService interface {
	Service[model.User, string]
	// 可以添加用户服务特有的方法
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

// userServiceImpl 用户服务实现
type userServiceImpl struct {
	*BaseService[model.User, string]
}

// NewUserService 创建用户服务
func NewUserService(baseURL string, opts ...DatasetOption) UserService {
	return &userServiceImpl{
		BaseService: NewBaseService[model.User, string](baseURL, opts...),
	}
}

// 可以添加用户服务特有的方法实现
func (s *userServiceImpl) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	result, err := s.List(ctx, 1, 1, map[string]any{
		"email": email,
	})

	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	return &result.Items[0], nil
}
