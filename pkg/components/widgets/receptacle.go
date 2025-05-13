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
	UpperActiveTab         string
	LowerActiveTab         string
	Document               string
	LayoutMode             string // "vertical" 或 "horizontal"
	IsUpperRibbonCollapsed bool
	IsLowerRibbonCollapsed bool
	upperMenu              types.RibbonMenu
	lowerMenu              types.RibbonMenu
	ErrorMessage           string
	isLoading              bool
	// HandleRibbonAction:func(buttonID string)
}

func (r *Receptacle) OnMount(ctx app.Context) {
	r.Document = "New Document"
	r.isLoading = true
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
			r.upperMenu = menu
			r.lowerMenu = menu
			r.isLoading = false
			// log.Printf("Ribbon menu loaded: %v", r.ribbonMenu)
		})

	})

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
func (r *Receptacle) OnUpperToggleCollapse(ctx app.Context) {
	log.Printf("Upper Ribbon collapsed state: %v", r.IsUpperRibbonCollapsed)
	r.IsUpperRibbonCollapsed = !r.IsUpperRibbonCollapsed
	log.Printf("Upper Ribbon collapsed state changed to: %v", r.IsUpperRibbonCollapsed)
	ctx.Update()
}
func (r *Receptacle) OnLowerToggleCollapse(ctx app.Context) {
	log.Printf("Lower Ribbon collapsed state: %v", r.IsLowerRibbonCollapsed)
	r.IsLowerRibbonCollapsed = !r.IsLowerRibbonCollapsed
	log.Printf("Lower Ribbon collapsed state changed to: %v", r.IsLowerRibbonCollapsed)
	ctx.Update()
}
func (re *Receptacle) onTabClick(ctx app.Context, tabID string) {
	re.UpperActiveTab = tabID
	log.Printf("receptacle Active tab changed to: %s", tabID)
}
func (re *Receptacle) Render() app.UI {
	// 根据布局模式应用不同的类
	layoutClass := "receptacle-vertical"
	if re.LayoutMode == "horizontal" {
		layoutClass = "receptacle-horizontal"
	}
	log.Printf("Receptacle upper activeTab: %s", re.UpperActiveTab)
	log.Printf("Receptacle lower activeTab: %s", re.LowerActiveTab)
	log.Printf("Receptacle upper collapsed: %v", re.IsUpperRibbonCollapsed)

	return app.Div().Class("receptacle", layoutClass).Body(
		&ribbon.Ribbon{
			RibbonMenu:  re.upperMenu,
			ActiveTab:   re.UpperActiveTab,
			IsCollapsed: re.IsUpperRibbonCollapsed,
			OnTabClick:  re.onTabClick,
			OnButtonClick: func(buttonID string) {
				re.handleRibbonAction(buttonID)
			},
			OnToggleCollapse: re.OnUpperToggleCollapse,
			IsLoading:        re.isLoading,
			ErrorMessage:     re.ErrorMessage,
			RibbonPosition:   "upper",
			LayoutMode:       re.LayoutMode,
		},
		&Workspace{ // Ensure the widgets package is correctly imported
			Document: re.Document,
			OnChange: func(document string) {
				re.Document = document
				// ctx.Update()
			},
		},
		&ribbon.Ribbon{
			RibbonMenu:  re.lowerMenu,
			ActiveTab:   re.LowerActiveTab,
			IsCollapsed: re.IsLowerRibbonCollapsed,
			OnTabClick: func(ctx app.Context, tabID string) {
				re.LowerActiveTab = tabID
				// ctx.Update()
			},
			OnButtonClick: func(buttonID string) {
				re.handleRibbonAction(buttonID)
			},
			OnToggleCollapse: re.OnLowerToggleCollapse,
			IsLoading:        re.isLoading,
			ErrorMessage:     re.ErrorMessage,
			RibbonPosition:   "lower",
			LayoutMode:       re.LayoutMode,
		},
	)

}
