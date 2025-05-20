package ribbon

// Ribbon组件

import (
	"github.com/gongzhaohui/gef/internal/frontend/components/ribbon/types"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TabBar 组件
type TabBar struct {
	app.Compo
	Tabs             []types.Tab
	ActiveTab        string
	IsCollapsed      bool
	OnTabClick       func(app.Context, string)
	LayoutMode       string
	RibbonPosition   string // "upper" 或 "lower"
	OnToggleCollapse func(app.Context)
}

func (t *TabBar) OnMount(ctx app.Context) {
}

func (t *TabBar) Render() app.UI {
	tabBarClass := "tab-bar"
	tabsClass := "tabs"
	actionsClass := "tab-bar-actions"
	buttonClass := "collapse-button"
	buttonText := "^"
	if t.IsCollapsed {
		buttonText = "v"
	}

	if t.LayoutMode == "horizontal" {
		tabBarClass += " horizontal"
		tabsClass += " horizontal"
		actionsClass += " horizontal"
		buttonClass += " horizontal"
		buttonText = "<"
		if t.IsCollapsed {
			buttonText = ">"
		}
	}

	tabs := make([]app.UI, len(t.Tabs))
	for i, tab := range t.Tabs {
		tabs[i] = &TabButton{
			Tab:        tab,
			ActiveTab:  t.ActiveTab,
			LayoutMode: t.LayoutMode,
			OnClick: func(ctx app.Context, e app.Event) {
				t.OnTabClick(ctx, tab.ID)
			},
		}
	}

	return app.Div().Class(tabBarClass).Body(app.Div().Class(tabsClass).Body(tabs...), app.Div().Class(actionsClass).Body(
		app.Button().Class(buttonClass).Text(buttonText).OnClick(func(ctx app.Context, e app.Event) {
			// ctx.NewAction("collapse")
			t.OnToggleCollapse(ctx)
		}),
	))

}

// TabButton 组件
type TabButton struct {
	app.Compo
	Tab       types.Tab
	ActiveTab string
	// IsCollapsed bool
	LayoutMode string
	OnClick    func(app.Context, app.Event)
}

func (t *TabButton) Render() app.UI {
	className := "tab"
	if t.Tab.ID == t.ActiveTab {
		className += " active"
	}

	if t.LayoutMode == "horizontal" {
		className += " horizontal"
	}
	return app.Button().Class(className).ID(t.Tab.ID).Text(t.Tab.Title).
		OnClick(func(ctx app.Context, e app.Event) {
			t.OnClick(ctx, e)
		})
}
