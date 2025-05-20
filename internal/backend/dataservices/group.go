package dataservices

import (
	"gorm.io/gorm"
)

// GroupService 组服务接口
type GroupService interface {
	GenericService
}

// GroupServiceImpl 组服务实现
type GroupServiceImpl struct {
	GenericService
}

// NewGroupService 创建组服务
func NewGroupService(db *gorm.DB) GroupService {
	return &GroupServiceImpl{
		GenericService: NewGenericService(db),
	}
}
