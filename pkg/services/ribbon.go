package services

import (
	"encoding/json"
	"gef/pkg/components/ribbon/types"
	"log"
	"net/http"
)

// RibbonService 提供Ribbon菜单相关服务
type RibbonService struct{}

// NewRibbonService 创建Ribbon服务实例
func NewRibbonService() *RibbonService {
	return &RibbonService{}
}

// LoadRibbonMenu 从服务器加载Ribbon菜单数据
func (s *RibbonService) LoadRibbonMenu() (types.RibbonMenu, error) {
	var menu types.RibbonMenu

	// 使用http包代替app.Get
	resp, err := http.Get("/web/ribbon_data.json")
	if err != nil {
		log.Printf("Failed to fetch ribbon data: %v", err)
		return menu, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("HTTP request failed with status code: %d", resp.StatusCode)
		return menu, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&menu); err != nil {
		log.Printf("Failed to parse ribbon data: %v", err)
		return menu, err
	}

	return menu, nil
}
