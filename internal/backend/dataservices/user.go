package dataservices

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gongzhaohui/gef/internal/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// JWT配置
type Claims struct {
	UserID      uint     `json:"user_id"`
	Username    string   `json:"username"`
	GroupNames  []string `json:"group_names"`
	RoleNames   []string `json:"role_names"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

const (
	secretKey      = "your-secret-key"       // JWT 签名密钥
	issuer         = "your-application-name" // 发行人
	tokenExpiresIn = 24 * time.Hour          // 令牌有效期
)

// UserService 用户服务接口
type UserService interface {
	GenericService
	GetUserByName(username string) (*models.User, error)
	Login(username, password string) (string, error)
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	GenericService
	db *gorm.DB
}

// NewUserService 创建用户服务
func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{
		GenericService: NewGenericService(db),
		db:             db,
	}
}

// GetUserByUsername 通过用户名获取用户
func (s *UserServiceImpl) GetUserByName(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Login 用户登录
func (s *UserServiceImpl) Login(username, password string) (string, error) {
	user, err := s.GetUserByName(username)
	if err != nil {
		return "", err
	}

	// 验证密码
	if err := comparePasswords(user.PasswordHash, password); err != nil {
		return "", err
	}

	// 生成令牌
	return generateToken(user.ID, username), nil
}

// Create 重写Create方法，处理密码哈希
func (s *UserServiceImpl) Create(entity interface{}) error {
	user, ok := entity.(*models.User)
	if !ok {
		return errors.New("实体类型错误")
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	return s.GenericService.Create(entity)
}

// comparePasswords 比较哈希密码和明文密码
func comparePasswords(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

// generateToken 生成JWT令牌
func generateToken(userID uint, username string) string {
	// 创建令牌声明
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiresIn).Unix(),
			Issuer:    issuer,
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}
