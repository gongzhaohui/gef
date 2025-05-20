package widgets

import (
	"log"

	"github.com/gongzhaohui/gef/internal/frontend/components/ribbon"
	"github.com/gongzhaohui/gef/internal/frontend/components/ribbon/types"
	"github.com/gongzhaohui/gef/internal/frontend/services"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Receptacle struct {
	app.Compo
	upperActiveTab   string
	lowerActiveTab   string
	Document         string
	LayoutMode       string // "vertical" 或 "horizontal"
	isUpperCollapsed bool
	isLowerCollapsed bool
	upperMenu        types.RibbonMenu
	lowerMenu        types.RibbonMenu
	ErrorMessage     string
	isLoading        bool
	// HandleRibbonAction:func(buttonID string)
}

func (re *Receptacle) OnInit() {
	re.isUpperCollapsed = false
	re.isLowerCollapsed = true
	re.upperActiveTab = "home"
	re.lowerActiveTab = "home"
	re.isLoading = false
	re.ErrorMessage = ""
}
func (re *Receptacle) OnMount(ctx app.Context) {
	re.Document = "New Document"
	re.isLowerCollapsed = true
	re.isUpperCollapsed = false
	re.lowerActiveTab = "home"
	re.upperActiveTab = "home"
	ctx.ObserveState("LowerActiveTab", &re.lowerActiveTab)

	// 使用服务加载Ribbon菜单数据
	ctx.Async(func() {
		re.isLoading = true
		ribbonService := services.NewRibbonService()
		menu, err := ribbonService.LoadRibbonMenu()
		if err != nil {
			re.ErrorMessage = "Failed to load ribbon data: " + err.Error()
			ctx.Update()
			return
		}
		ctx.Dispatch(func(ctx app.Context) {
			re.upperMenu = menu
			re.lowerMenu = menu
			re.isLoading = false
			// log.Printf("Ribbon menu loaded: %v", menu)
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

	default:
		// Handle other actions
		r.ErrorMessage = "Unknown action: " + buttonID
	}
}
func (r *Receptacle) OnUpperToggleCollapse(ctx app.Context) {
	r.isUpperCollapsed = !r.isUpperCollapsed
}
func (r *Receptacle) OnLowerToggleCollapse(ctx app.Context) {
	r.isLowerCollapsed = !r.isLowerCollapsed
}
func (re *Receptacle) onUpperTabClick(ctx app.Context, tabID string) {
	re.upperActiveTab = tabID
}
func (re *Receptacle) onLowerTabClick(ctx app.Context, tabID string) {
	re.lowerActiveTab = tabID
}
func (re *Receptacle) Render() app.UI {

	// 根据布局模式应用不同的类
	layoutClass := "receptacle-vertical"
	if re.LayoutMode == "horizontal" {
		layoutClass = "receptacle-horizontal"
	}
	return app.Div().Class("receptacle", layoutClass).Body(
		&ribbon.Ribbon{
			RibbonMenu:  re.upperMenu,
			ActiveTab:   re.upperActiveTab,
			IsCollapsed: re.isUpperCollapsed,
			OnTabClick:  re.onUpperTabClick,
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
			ActiveTab:   re.lowerActiveTab,
			IsCollapsed: re.isLowerCollapsed,
			OnTabClick:  re.onLowerTabClick,
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
