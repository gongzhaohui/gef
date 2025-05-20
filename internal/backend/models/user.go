package models

import (
	"gorm.io/gorm"
)

type User_Dto struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       int    `json:"age,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// User 用户模型
type User struct {
	gorm.Model
	Username     string      `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string      `gorm:"not null" json:"-"`
	UserGroups   []UserGroup `gorm:"foreignKey:UserID" json:"groups,omitempty"`
}

// Group 组模型
type Group struct {
	gorm.Model
	Name        string      `gorm:"uniqueIndex;not null" json:"name"`
	Description string      `gorm:"type:text" json:"description,omitempty"`
	GroupRoles  []GroupRole `gorm:"foreignKey:GroupID" json:"roles,omitempty"`
}

// Role 角色模型
type Role struct {
	gorm.Model
	Name            string           `gorm:"uniqueIndex;not null" json:"name"`
	Description     string           `gorm:"type:text" json:"description,omitempty"`
	RolePermissions []RolePermission `gorm:"foreignKey:RoleID" json:"permissions,omitempty"`
}

// Permission 权限模型
type Permission struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}

// UserGroup 用户-组关联模型
type UserGroup struct {
	UserID  uint `gorm:"primaryKey" json:"user_id"`
	GroupID uint `gorm:"primaryKey" json:"group_id"`
}

// GroupRole 组-角色关联模型
type GroupRole struct {
	GroupID uint `gorm:"primaryKey" json:"group_id"`
	RoleID  uint `gorm:"primaryKey" json:"role_id"`
}

// RolePermission 角色-权限关联模型
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey" json:"role_id"`
	PermissionID uint `gorm:"primaryKey" json:"permission_id"`
}
