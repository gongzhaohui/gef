package dataservices

import "gorm.io/gorm"

// RoleService 角色服务接口
type RoleService interface {
	GenericService
}

// RoleServiceImpl 角色服务实现
type RoleServiceImpl struct {
	GenericService
}

// NewRoleService 创建角色服务
func NewRoleService(db *gorm.DB) RoleService {
	return &RoleServiceImpl{
		GenericService: NewGenericService(db),
	}
}
