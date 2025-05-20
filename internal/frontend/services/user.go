package services

import (
	"context"

	"github.com/gongzhaohui/gef/internal/backend/models" // 假设 model 包已正确导入
)

// UserService 用户服务接口
type UserService interface {
	Service[models.User_Dto, string]
	// 可以添加用户服务特有的方法
	GetByEmail(ctx context.Context, email string) (*models.User_Dto, error)
}

// userServiceImpl 用户服务实现
type userServiceImpl struct {
	*BaseService[models.User_Dto, string]
}

// NewUserService 创建用户服务
func NewUserService(baseURL string, opts ...DatasetOption) UserService {
	return &userServiceImpl{
		BaseService: NewBaseService[models.User_Dto, string](baseURL, opts...),
	}
}

// 可以添加用户服务特有的方法实现
func (s *userServiceImpl) GetByEmail(ctx context.Context, email string) (*models.User_Dto, error) {
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
