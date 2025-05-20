package dataservices

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// BaseService 基础服务接口
type BaseService interface {
	Create(entity interface{}) error
	GetByID(id uint, entity interface{}) error
	Update(entity interface{}) error
	Delete(id uint, entity interface{}) error
	GetAll(entity interface{}) error
}

// BaseServiceImpl 基础服务实现
type BaseServiceImpl struct {
	db *gorm.DB
}

// NewBaseService 创建基础服务
func NewBaseService(db *gorm.DB) BaseService {
	return &BaseServiceImpl{db: db}
}

// Create 创建实体
func (s *BaseServiceImpl) Create(entity interface{}) error {
	return s.db.Create(entity).Error
}

// GetByID 通过ID获取实体
func (s *BaseServiceImpl) GetByID(id uint, entity interface{}) error {
	return s.db.First(entity, id).Error
}

// Update 更新实体
func (s *BaseServiceImpl) Update(entity interface{}) error {
	return s.db.Save(entity).Error
}

// Delete 删除实体
func (s *BaseServiceImpl) Delete(id uint, entity interface{}) error {
	// 获取表名
	tableName := s.db.Model(entity).Statement.Table

	// 执行删除
	return s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName), id).Error
}

// GetAll 获取所有实体
func (s *BaseServiceImpl) GetAll(entity interface{}) error {
	return s.db.Find(entity).Error
}

// GenericService 通用服务接口，继承BaseService并添加关联管理方法
type GenericService interface {
	BaseService
	AddRelation(parentID, childID uint, relation interface{}) error
	RemoveRelation(parentID, childID uint, relation interface{}) error
	GetRelations(parentID uint, relations interface{}) error
}

// GenericServiceImpl 通用服务实现
type GenericServiceImpl struct {
	BaseService
	db *gorm.DB
}

// NewGenericService 创建通用服务
func NewGenericService(db *gorm.DB) GenericService {
	return &GenericServiceImpl{
		BaseService: NewBaseService(db),
		db:          db,
	}
}

// AddRelation 添加关联
func (s *GenericServiceImpl) AddRelation(parentID, childID uint, relation interface{}) error {
	// 检查关联是否已存在
	var count int64
	s.db.Model(relation).Where("parent_id = ? AND child_id = ?", parentID, childID).Count(&count)
	if count > 0 {
		return errors.New("关联已存在")
	}

	// 设置关联ID
	setFieldValue(relation, "ParentID", parentID)
	setFieldValue(relation, "ChildID", childID)

	return s.db.Create(relation).Error
}

// RemoveRelation 移除关联
func (s *GenericServiceImpl) RemoveRelation(parentID, childID uint, relation interface{}) error {
	// 获取表名
	// 获取表名
	tableName := s.db.Model(relation).Statement.Table

	// 执行删除
	return s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE parent_id = ? AND child_id = ?", tableName), parentID, childID).Error
}

// GetRelations 获取关联
func (s *GenericServiceImpl) GetRelations(parentID uint, relations interface{}) error {
	// 获取表名
	tableName := s.db.Model(relations).Statement.Table

	// 执行查询
	return s.db.Raw(fmt.Sprintf("SELECT * FROM %s WHERE parent_id = ?", tableName), parentID).Scan(relations).Error
}

// setFieldValue 设置结构体字段值
func setFieldValue(obj interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(obj).Elem()
	if f := v.FieldByName(fieldName); f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(value))
	}
}
