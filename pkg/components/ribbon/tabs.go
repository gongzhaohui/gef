package ribbon

// Ribbon组件

import (
	"gef/pkg/components/ribbon/types"

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
	if t.IsCollapsed {
	}
	if t.LayoutMode == "horizontal" {
		tabBarClass += " horizontal"
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

	return app.Div().Class(tabBarClass).Body(app.Div().Class("tabs").Body(tabs...), app.Div().Class("tab-bar-actions").Body(
		app.Button().Class("collapse-button").Text("Collapse").OnClick(func(ctx app.Context, e app.Event) {
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
