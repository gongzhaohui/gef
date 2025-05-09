package ribbon

// Ribbon组件

import (
	"gef/pkg/components/ribbon/types"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TabBar 组件
type TabBar struct {
	app.Compo
	Tabs        []types.Tab
	ActiveTab   string
	IsCollapsed bool
	OnClick     func(string)
	LayoutMode  string
}

func (t *TabBar) OnMount(ctx app.Context) {
	ctx.Handle("toggleLayout", t.ToggleLayout)
}

// ToggleLayout toggles the layout mode between "horizontal" and "vertical".
func (t *TabBar) ToggleLayout(ctx app.Context, a app.Action) {
	if t.LayoutMode == "horizontal" {
		t.LayoutMode = "vertical"
	} else {
		t.LayoutMode = "horizontal"
	}
	ctx.Update()
}
func (t *TabBar) Render() app.UI {
	tabBarClass := "tab-bar"
	if t.LayoutMode == "horizontal" {
		tabBarClass += " horizontal"
	}

	tabs := make([]app.UI, len(t.Tabs))
	for i, tab := range t.Tabs {
		tabs[i] = &TabButton{
			Tab:         tab,
			ActiveTab:   t.ActiveTab,
			IsCollapsed: t.IsCollapsed,
			LayoutMode:  t.LayoutMode,
			OnClick:     func() { t.OnClick(tab.ID) },
		}
	}

	return app.Div().Class(tabBarClass).Body(tabs...)
}

// TabButton 组件
type TabButton struct {
	app.Compo
	Tab         types.Tab
	ActiveTab   string
	IsCollapsed bool
	LayoutMode  string
	OnClick     func()
}

func (t *TabButton) Render() app.UI {
	className := "tab"
	if t.Tab.ID == t.ActiveTab {
		className += " active"
	}
	if t.IsCollapsed {
		className += " collapsed"
	}
	if t.LayoutMode == "horizontal" {
		className += " horizontal"
	}
	return app.Button().Class(className).ID(t.Tab.ID).Text(t.Tab.Title).
		OnClick(func(ctx app.Context, e app.Event) {
			t.OnClick()
		})
}
