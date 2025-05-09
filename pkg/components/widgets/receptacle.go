package widgets

import (
	"gef/pkg/components/ribbon"
	"gef/pkg/components/ribbon/types"
	"gef/pkg/services"
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Receptacle struct {
	app.Compo
	activeTab              string
	Document               string
	LayoutMode             string // "vertical" 或 "horizontal"
	IsUpperRibbonCollapsed bool
	IsLowerRibbonCollapsed bool
	ribbonMenu             types.RibbonMenu
	ErrorMessage           string
	isLoading              bool
	// HandleRibbonAction:func(buttonID string)
}

func (r *Receptacle) OnMount(ctx app.Context) {
	r.activeTab = "home"
	r.Document = "New Document"
	// r.LayoutMode = "vertical" // 默认垂直布局（上中下）
	r.IsUpperRibbonCollapsed = false
	r.IsLowerRibbonCollapsed = true
	// 初始化Ribbon组件
	// r.LayoutMode = "vertical" // 默认垂直布局（上中下）
	r.isLoading = true
	log.Printf("Ribbon component mounted with initial state: %v", r)
	ctx.Handle("toggleLayout", r.ToggleLayout)
	// UpdateLayout updates the layout mode of the Receptacle.
	log.Printf("Receptacle layout mode : %v", r.LayoutMode)
	// 使用服务加载Ribbon菜单数据
	ctx.Async(func() {
		ribbonService := services.NewRibbonService()
		menu, err := ribbonService.LoadRibbonMenu()
		if err != nil {
			r.ErrorMessage = "Failed to load ribbon data: " + err.Error()
			ctx.Update()
			return
		}
		ctx.Dispatch(func(ctx app.Context) {
			r.ribbonMenu = menu
			r.isLoading = false
			// log.Printf("Ribbon menu loaded: %v", r.ribbonMenu)
		})

	})

}
func (r *Receptacle) ToggleLayout(ctx app.Context, action app.Action) {

	if r.LayoutMode == "vertical" {
		r.LayoutMode = "horizontal"
	} else {
		r.LayoutMode = "vertical"
	}
	log.Printf("Receptacle layout mode changed to: %v", r.LayoutMode)
	ctx.Update()
}
func (r *Receptacle) handleRibbonAction(buttonID string) {
	log.Printf("Ribbon action: %s", buttonID)
	// Add logic to handle ribbon button actions
	switch buttonID {
	case "save":
		// Example: Save the document
		r.Document = "Document Saved"
	case "open":
		// Example: Open a new document
		r.Document = "New Document Opened"
	// case "toggle-layout":
	// 	if r.LayoutMode == "vertical" {
	// 		r.LayoutMode = "horizontal"
	// 	} else {
	// 		r.LayoutMode = "vertical"
	// 	}
	// ctx.Update()

	default:
		// Handle other actions
		r.ErrorMessage = "Unknown action: " + buttonID
	}
}
func (r *Receptacle) Render() app.UI {
	// 根据布局模式应用不同的类
	layoutClass := "receptacle-vertical"
	if r.LayoutMode == "horizontal" {
		layoutClass = "receptacle-horizontal"
	}

	return app.Div().Class("receptacle", layoutClass).Body(
		&ribbon.Ribbon{
			RibbonMenu:  r.ribbonMenu,
			ActiveTab:   r.activeTab,
			IsCollapsed: false,
			OnTabClick: func(tabID string) {
				r.activeTab = tabID
				r.IsUpperRibbonCollapsed = false
				log.Printf("Active tab changed to: %s", tabID)
				// ctx.Update()
			},
			OnButtonClick: func(buttonID string) {
				r.handleRibbonAction(buttonID)
			},
			OnToggleCollapse: func() {
				r.IsLowerRibbonCollapsed = !r.IsLowerRibbonCollapsed
				// ctx.Update()
			},
			IsLoading:      r.isLoading,
			ErrorMessage:   r.ErrorMessage,
			RibbonPosition: "upper",
			LayoutMode:     r.LayoutMode,
		},
		&Workspace{ // Ensure the widgets package is correctly imported
			Document: r.Document,
			OnChange: func(document string) {
				r.Document = document
				// ctx.Update()
			},
		},
		&ribbon.Ribbon{
			RibbonMenu:  r.ribbonMenu,
			ActiveTab:   r.activeTab,
			IsCollapsed: true,
			OnTabClick: func(tabID string) {
				r.activeTab = tabID
				// ctx.Update()
			},
			OnButtonClick: func(buttonID string) {
				r.handleRibbonAction(buttonID)
			},
			OnToggleCollapse: func() {
				r.IsLowerRibbonCollapsed = !r.IsLowerRibbonCollapsed
				// ctx.Update()
			},
			IsLoading:      r.isLoading,
			ErrorMessage:   r.ErrorMessage,
			RibbonPosition: "lower",
			LayoutMode:     r.LayoutMode,
		},
	)

}
