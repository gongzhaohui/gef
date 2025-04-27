package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	// 可注入数据库等依赖
}

func (h *OrderHandler) GetOrders(c echo.Context) error {
	// 模拟数据
	orders := []Order{
		{ID: "1001", Amount: 2990, Status: "已完成"},
		{ID: "1002", Amount: 1500, Status: "处理中"},
	}

	return c.JSON(http.StatusOK, orders)
}

// 数据模型
type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}
