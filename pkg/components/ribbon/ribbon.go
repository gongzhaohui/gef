package ribbon

import (
	"gef/pkg/services/types"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Ribbon组件
type Ribbon struct {
	app.Compo
	RibbonMenu       types.RibbonMenu
	ActiveTab        string
	RibbonLayout     string
	IsDragging       bool
	IsCollapsed      bool
	OnTabClick       func(string)
	OnDragStart      func(int, int)
	OnDragMove       func(int, int)
	OnDragEnd        func()
	OnButtonClick    func(string)
	OnToggleCollapse func()
	IsLoading        bool
	ErrorMessage     string
}

func (r *Ribbon) Render() app.UI {
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
	if r.RibbonLayout == "vertical" {
		ribbonClass += " vertical"
	}
	if r.IsCollapsed {
		ribbonClass += " collapsed"
	}

	return app.Div().Class(ribbonClass).Body(
		&FileMenu{
			Items: []string{"New", "Open", "Save", "Exit"},
		},
		&TabBar{
			Tabs:         r.RibbonMenu.Tabs,
			ActiveTab:    r.ActiveTab,
			RibbonLayout: r.RibbonLayout,
			IsCollapsed:  r.IsCollapsed,
			OnClick:      r.OnTabClick,
		},
		app.If(!r.IsCollapsed, func() app.UI {
			return app.Div().Class("ribbon-content").Body(
				&DragHandle{
					RibbonLayout: r.RibbonLayout,
					OnDragStart:  r.OnDragStart,
					OnDragMove:   r.OnDragMove,
					OnDragEnd:    r.OnDragEnd,
				},
				&TabContent{
					Tabs:          r.RibbonMenu.Tabs,
					ActiveTab:     r.ActiveTab,
					RibbonLayout:  r.RibbonLayout,
					OnButtonClick: r.OnButtonClick,
				},
			)
		}),
		&RibbonFooter{
			IsCollapsed:      r.IsCollapsed,
			OnToggleCollapse: r.OnToggleCollapse,
		},
	)
}

// 标签栏组件
type TabBar struct {
	app.Compo
	Tabs         []types.Tab
	ActiveTab    string
	RibbonLayout string
	IsCollapsed  bool
	OnClick      func(string)
}

func (t *TabBar) Render() app.UI {
	tabBarClass := "tab-bar"
	if t.RibbonLayout == "vertical" {
		tabBarClass += " vertical"
	}

	tabs := make([]app.UI, len(t.Tabs))
	for i, tab := range t.Tabs {
		tabs[i] = &TabButton{
			Tab:         tab,
			ActiveTab:   t.ActiveTab,
			IsCollapsed: t.IsCollapsed,
			OnClick:     func() { t.OnClick(tab.ID) },
		}
	}

	return app.Div().Class(tabBarClass).Body(tabs...)
}

// 标签按钮组件
type TabButton struct {
	app.Compo
	Tab         types.Tab
	ActiveTab   string
	IsCollapsed bool
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

	return app.Button().Class(className).ID(t.Tab.ID).Text(t.Tab.Title).
		OnClick(func(ctx app.Context, e app.Event) {
			t.OnClick()
		})
}

// Ribbon底部组件
type RibbonFooter struct {
	app.Compo
	IsCollapsed      bool
	OnToggleCollapse func()
}

func (r *RibbonFooter) Render() app.UI {
	buttonClass := "toggle-button"
	buttonIcon := "fa-chevron-up"
	buttonTitle := "Collapse Ribbon"

	if r.IsCollapsed {
		buttonClass += " collapsed"
		buttonIcon = "fa-chevron-down"
		buttonTitle = "Expand Ribbon"
	}

	return app.Div().Class("ribbon-footer").Body(
		app.Button().Class(buttonClass).Title(buttonTitle).Body(
			app.I().Class("fa", buttonIcon),
		).OnClick(func(ctx app.Context, e app.Event) {
			r.OnToggleCollapse()
		}),
	)
}

// FileMenu组件
type FileMenu struct {
	app.Compo
	Items []string
}

func (f *FileMenu) Render() app.UI {
	items := make([]app.UI, len(f.Items))
	for i, item := range f.Items {
		items[i] = app.Div().Class("file-menu-item").Text(item)
	}

	return app.Div().Class("file-menu").Body(items...)
}

// DragHandle组件
type DragHandle struct {
	app.Compo
	RibbonLayout string
	OnDragStart  func(int, int)
	OnDragMove   func(int, int)
	OnDragEnd    func()
}

func (d *DragHandle) Render() app.UI {
	className := "drag-handle"
	if d.RibbonLayout == "vertical" {
		className += " vertical"
	} else {
		className += " horizontal"
	}

	return app.Div().Class(className).Body(
		app.Div().Class("drag-area").OnMouseDown(func(ctx app.Context, e app.Event) {
			if d.OnDragStart != nil {
				d.OnDragStart(e.Get("clientX").Int(), e.Get("clientY").Int())
			}
		}).OnMouseMove(func(ctx app.Context, e app.Event) {
			if d.OnDragMove != nil {
				d.OnDragMove(e.Get("clientX").Int(), e.Get("clientY").Int())
			}
		}).OnMouseUp(func(ctx app.Context, e app.Event) {
			if d.OnDragEnd != nil {
				d.OnDragEnd()
			}
		}),
	)
}

// TabContent组件
type TabContent struct {
	app.Compo
	Tabs          []types.Tab
	ActiveTab     string
	RibbonLayout  string
	OnButtonClick func(string)
}

func (t *TabContent) Render() app.UI {
	content := make([]app.UI, len(t.Tabs))
	for i, tab := range t.Tabs {
		if tab.ID == t.ActiveTab {
			content[i] = app.Div().Class("tab-content").Body(
				app.Div().Class("tab-title").Text(tab.Title),
				app.Range(tab.Groups).Slice(func(j int) app.UI {
					group := tab.Groups[j]
					return app.Div().Class("tab-group").Body(
						app.Range(group.Buttons).Slice(func(i int) app.UI {
							button := group.Buttons[i]
							return app.Button().Class("tab-button").Text(button.Title).
								OnClick(func(ctx app.Context, e app.Event) {
									t.OnButtonClick(button.ID)
								})

						}),
					)
				},
				),
			)
		}
	}

	return app.Div().Class("tab-content-container").Body(content...)
}

// 其他Ribbon组件保持不变...

// 其他Ribbon组件保持不变...
