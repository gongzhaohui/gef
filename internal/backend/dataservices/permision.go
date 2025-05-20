package dataservices

import "gorm.io/gorm"

// PermissionService 权限服务接口
type PermissionService interface {
	GenericService
}

// PermissionServiceImpl 权限服务实现
type PermissionServiceImpl struct {
	GenericService
}

// NewPermissionService 创建权限服务
func NewPermissionService(db *gorm.DB) PermissionService {
	return &PermissionServiceImpl{
		GenericService: NewGenericService(db),
	}
}
