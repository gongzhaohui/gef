package controllers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gongzhaohui/gef/internal/backend/dataservices"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BaseController 基础控制器
type BaseController struct {
	service    dataservices.BaseService
	entityType reflect.Type
}

// NewBaseController 创建基础控制器
func NewBaseController(service dataservices.BaseService, entityType reflect.Type) *BaseController {
	return &BaseController{
		service:    service,
		entityType: entityType,
	}
}

// Create 创建实体
func (c *BaseController) Create(ctx echo.Context) error {
	// 创建新的实体实例
	entity := reflect.New(c.entityType).Elem().Interface()

	// 绑定请求数据
	if err := ctx.Bind(entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求格式"})
	}

	// 验证数据
	if err := ctx.Validate(entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 创建实体
	if err := c.service.Create(entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "创建失败"})
	}

	return ctx.JSON(http.StatusCreated, entity)
}

// GetByID 通过ID获取实体
func (c *BaseController) GetByID(ctx echo.Context) error {
	// 获取ID参数
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	// 创建新的实体实例
	entity := reflect.New(c.entityType).Elem().Interface()

	// 获取实体
	if err := c.service.GetByID(uint(id), entity); err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "实体不存在"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "获取失败"})
	}

	return ctx.JSON(http.StatusOK, entity)
}

// Update 更新实体
func (c *BaseController) Update(ctx echo.Context) error {
	// 获取ID参数
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	// 创建新的实体实例
	entity := reflect.New(c.entityType).Elem().Interface()

	// 设置ID
	setID(entity, uint(id))

	// 绑定请求数据
	if err := ctx.Bind(entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求格式"})
	}

	// 验证数据
	if err := ctx.Validate(entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 更新实体
	if err := c.service.Update(entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "更新失败"})
	}

	return ctx.JSON(http.StatusOK, entity)
}

// Delete 删除实体
func (c *BaseController) Delete(ctx echo.Context) error {
	// 获取ID参数
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	// 创建新的实体实例
	entity := reflect.New(c.entityType).Elem().Interface()

	// 删除实体
	if err := c.service.Delete(uint(id), entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "删除失败"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "删除成功"})
}

// GetAll 获取所有实体
func (c *BaseController) GetAll(ctx echo.Context) error {
	// 创建切片实例
	sliceType := reflect.SliceOf(c.entityType)
	entities := reflect.New(sliceType).Elem().Interface()

	// 获取所有实体
	if err := c.service.GetAll(entities); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "获取失败"})
	}

	return ctx.JSON(http.StatusOK, entities)
}

// GenericController 通用控制器，继承BaseController并添加关联管理方法
type GenericController struct {
	*BaseController
	service      dataservices.GenericService
	relationType reflect.Type
}

// NewGenericController 创建通用控制器
func NewGenericController(service dataservices.GenericService, entityType, relationType reflect.Type) *GenericController {
	return &GenericController{
		BaseController: NewBaseController(service, entityType),
		service:        service,
		relationType:   relationType,
	}
}

// AddRelation 添加关联
func (c *GenericController) AddRelation(ctx echo.Context) error {
	// 获取ID参数
	parentID, err := strconv.Atoi(ctx.Param("parent_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的父ID"})
	}

	childID, err := strconv.Atoi(ctx.Param("child_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的子ID"})
	}

	// 创建新的关联实例
	relation := reflect.New(c.relationType).Elem().Interface()

	// 添加关联
	if err := c.service.AddRelation(uint(parentID), uint(childID), relation); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "关联已添加"})
}

// RemoveRelation 移除关联
func (c *GenericController) RemoveRelation(ctx echo.Context) error {
	// 获取ID参数
	parentID, err := strconv.Atoi(ctx.Param("parent_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的父ID"})
	}

	childID, err := strconv.Atoi(ctx.Param("child_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的子ID"})
	}

	// 创建新的关联实例
	relation := reflect.New(c.relationType).Elem().Interface()

	// 移除关联
	if err := c.service.RemoveRelation(uint(parentID), uint(childID), relation); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "关联已移除"})
}

// GetRelations 获取关联
func (c *GenericController) GetRelations(ctx echo.Context) error {
	// 获取ID参数
	parentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "无效的ID"})
	}

	// 创建切片实例
	sliceType := reflect.SliceOf(c.relationType)
	relations := reflect.New(sliceType).Elem().Interface()

	// 获取关联
	if err := c.service.GetRelations(uint(parentID), relations); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "获取关联失败"})
	}

	return ctx.JSON(http.StatusOK, relations)
}

// setID 设置结构体的ID字段
func setID(obj interface{}, id uint) {
	v := reflect.ValueOf(obj).Elem()
	if f := v.FieldByName("ID"); f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(id))
	}
}
