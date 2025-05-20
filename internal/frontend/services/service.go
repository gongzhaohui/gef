package services

import (
	"context"
	"fmt"
)

// Service 通用服务接口
type Service[T any, ID comparable] interface {
	List(ctx context.Context, page, pageSize int, filter map[string]any) (*PagedResult[T], error)
	Get(ctx context.Context, id ID) (*T, error)
	Create(ctx context.Context, item T) (*T, error)
	Update(ctx context.Context, id ID, item T) (*T, error)
	Delete(ctx context.Context, id ID) error
}

// BaseService 基础服务实现
type BaseService[T any, ID comparable] struct {
	dataset *Dataset[T]
}

// NewBaseService 创建基础服务
func NewBaseService[T any, ID comparable](baseURL string, opts ...DatasetOption) *BaseService[T, ID] {
	return &BaseService[T, ID]{
		dataset: NewDataset[T](baseURL, opts...),
	}
}

// List 实现 Service 接口的 List 方法
func (s *BaseService[T, ID]) List(ctx context.Context, page, pageSize int, filter map[string]any) (*PagedResult[T], error) {
	return s.dataset.List(ctx, page, pageSize, filter)
}

// Get 实现 Service 接口的 Get 方法
func (s *BaseService[T, ID]) Get(ctx context.Context, id ID) (*T, error) {
	return s.dataset.Get(ctx, fmt.Sprintf("%v", id))
}

// Create 实现 Service 接口的 Create 方法
func (s *BaseService[T, ID]) Create(ctx context.Context, item T) (*T, error) {
	return s.dataset.Create(ctx, item)
}

// Update 实现 Service 接口的 Update 方法
func (s *BaseService[T, ID]) Update(ctx context.Context, id ID, item T) (*T, error) {
	return s.dataset.Update(ctx, fmt.Sprintf("%v", id), item)
}

// Delete 实现 Service 接口的 Delete 方法
func (s *BaseService[T, ID]) Delete(ctx context.Context, id ID) error {
	return s.dataset.Delete(ctx, fmt.Sprintf("%v", id))
}
