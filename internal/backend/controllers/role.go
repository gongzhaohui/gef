package controllers

import (
	"reflect"

	"github.com/gongzhaohui/gef/internal/backend/dataservices"
	"github.com/gongzhaohui/gef/internal/backend/models"
)

// RoleController 角色控制器
type RoleController struct {
	*GenericController
}

// NewRoleController 创建角色控制器
func NewRoleController(roleService dataservices.RoleService) *RoleController {
	return &RoleController{
		GenericController: NewGenericController(
			roleService,
			reflect.TypeOf(models.Role{}),
			reflect.TypeOf(models.RolePermission{}),
		),
	}
}
