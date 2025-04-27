package main

import (
	"gef/pkg/api"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// 中间件
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// 注册路由
	handler := &api.OrderHandler{}
	e.GET("/api/orders", handler.GetOrders)
	rootDir, _ := os.Getwd()
	webDir := filepath.Join(rootDir, "../../web")
	e.Logger.Print("webDir:", webDir)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasSuffix(c.Request().URL.Path, ".wasm") {
				c.Response().Header().Set("Content-Type", "application/wasm")
			}
			return next(c)
		}
	})
	// 静态文件服务（开发模式）
	e.Static("/", webDir)
	e.Static("/web", webDir)

	e.Logger.Fatal(e.Start(":3000"))
}
