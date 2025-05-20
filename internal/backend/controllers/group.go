package controllers

import (
	"reflect"

	"github.com/gongzhaohui/gef/internal/backend/dataservices"
	"github.com/gongzhaohui/gef/internal/backend/models"
)

// GroupController 组控制器
type GroupController struct {
	*GenericController
}

// NewGroupController 创建组控制器
func NewGroupController(groupService dataservices.GroupService) *GroupController {
	return &GroupController{
		GenericController: NewGenericController(
			groupService,
			reflect.TypeOf(models.Group{}),
			reflect.TypeOf(models.GroupRole{}),
		),
	}
}
