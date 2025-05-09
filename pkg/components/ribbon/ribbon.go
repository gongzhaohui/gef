package ribbon

import (
	"gef/pkg/components/ribbon/types"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Ribbon组件
type Ribbon struct {
	app.Compo
	RibbonMenu       types.RibbonMenu
	ActiveTab        string
	IsCollapsed      bool
	OnTabClick       func(string)
	OnButtonClick    func(string)
	OnToggleCollapse func()
	IsLoading        bool
	ErrorMessage     string
	RibbonPosition   string // "upper" 或 "lower"
	LayoutMode       string // "vertical" 或 "horizontal"
}

func (r *Ribbon) OnMount(ctx app.Context) {
	ctx.Handle("toggleLayout", r.ToggleLayout)
}

// ToggleLayout toggles the layout mode between "vertical" and "horizontal".
func (r *Ribbon) ToggleLayout(ctx app.Context, action app.Action) {
	if r.LayoutMode == "vertical" {
		r.LayoutMode = "horizontal"
	} else {
		r.LayoutMode = "vertical"
	}
	ctx.Update()
}
func (r *Ribbon) Render() app.UI {
	// log.Printf("Ribbon component rendering: %v", r.RibbonMenu)
	if r.IsLoading {
		return app.Div().Class("ribbon").Body(
			app.Div().Class("loading-indicator").Text("Loading Ribbon..."),
		)
	}

	if r.ErrorMessage != "" {
		return app.Div().Class("ribbon").Body(
			app.Div().Class("error-message").Text(r.ErrorMessage),
		)
	}

	ribbonClass := "ribbon"
	if r.LayoutMode == "horizontal" {
		ribbonClass += " horizontal"
	}

	if r.RibbonPosition == "lower" {
		ribbonClass += " lower-ribbon"
	}

	if r.IsCollapsed {
		ribbonClass += " collapsed"
	}

	return app.Div().Class(ribbonClass).Body(

		&TabBar{
			Tabs:        r.RibbonMenu.Tabs,
			ActiveTab:   r.ActiveTab,
			IsCollapsed: r.IsCollapsed,
			OnClick:     r.OnTabClick,
			LayoutMode:  r.LayoutMode,
		},
		app.If(!r.IsCollapsed, func() app.UI {
			return &TabContent{
				Tabs:          r.RibbonMenu.Tabs,
				ActiveTab:     r.ActiveTab,
				LayoutMode:    r.LayoutMode,
				OnButtonClick: r.OnButtonClick,
			}
		}),
		&RibbonFooter{
			IsCollapsed:      r.IsCollapsed,
			OnToggleCollapse: r.OnToggleCollapse,
			RibbonPosition:   r.RibbonPosition,
		},
	)
}
