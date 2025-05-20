package controllers

import (
	"net/http"
	"reflect"

	"github.com/gongzhaohui/gef/internal/backend/dataservices"
	"github.com/gongzhaohui/gef/internal/backend/models"
	"github.com/labstack/echo/v4"
)

// UserController 用户控制器
type UserController struct {
	*GenericController
	userService dataservices.UserService
}

// NewUserController 创建用户控制器
func NewUserController(userService dataservices.UserService) *UserController {
	return &UserController{
		GenericController: NewGenericController(
			userService,
			reflect.TypeOf(models.User{}),
			reflect.TypeOf(models.UserGroup{}),
		),
		userService: userService,
	}
}

// Login 用户登录
func (c *UserController) Login(ctx echo.Context) error {
	var req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// 绑定请求数据
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求格式"})
	}

	// 验证数据
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 登录
	token, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "用户名或密码错误"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
