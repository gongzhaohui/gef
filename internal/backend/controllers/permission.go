package controllers

import (
	"reflect"

	"github.com/gongzhaohui/gef/internal/backend/dataservices"
	"github.com/gongzhaohui/gef/internal/backend/models"
)

// PermissionController 权限控制器
type PermissionController struct {
	*GenericController
}

// NewPermissionController 创建权限控制器
func NewPermissionController(permissionService dataservices.PermissionService) *PermissionController {
	return &PermissionController{
		GenericController: NewGenericController(
			permissionService,
			reflect.TypeOf(models.Permission{}),
			reflect.TypeOf(models.RolePermission{}),
		),
	}
}
