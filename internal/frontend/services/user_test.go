package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gongzhaohui/gef/internal/backend/models"
)

func test() {
	ctx := context.Background()

	// 创建用户服务
	userService := NewUserService(
		"https://api.example.com/v1/users",
		WithTimeout(10*time.Second),
		WithHeaders(map[string]string{
			"X-App-Version": "1.0.0",
		}),
		WithAuthToken("your-auth-token"),
	)

	// 创建用户
	newUser := models.User_Dto{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	createdUser, err := userService.Create(ctx, newUser)
	if err != nil {
		log.Fatalf("创建用户失败: %v", err)
	}

	fmt.Printf("创建的用户: %+v\n", createdUser)

	// 获取用户列表
	users, err := userService.List(ctx, 1, 10, map[string]any{
		"age": map[string]int{
			"gt": 25,
		},
	})
	if err != nil {
		log.Fatalf("获取用户列表失败: %v", err)
	}

	fmt.Printf("获取到 %d 个用户，总共有 %d 个用户\n", len(users.Items), users.Total)

	// 通过邮箱获取用户
	userByEmail, err := userService.GetByEmail(ctx, "john@example.com")
	if err != nil {
		log.Fatalf("通过邮箱获取用户失败: %v", err)
	}

	if userByEmail != nil {
		fmt.Printf("通过邮箱获取的用户: %+v\n", userByEmail)
	}

	// 更新用户
	createdUser.Age = 31
	updatedUser, err := userService.Update(ctx, createdUser.ID, *createdUser)
	if err != nil {
		log.Fatalf("更新用户失败: %v", err)
	}

	fmt.Printf("更新后的用户: %+v\n", updatedUser)

	// 删除用户
	if err := userService.Delete(ctx, createdUser.ID); err != nil {
		log.Fatalf("删除用户失败: %v", err)
	}

	fmt.Println("用户已成功删除")
}
