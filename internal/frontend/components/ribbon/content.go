package ribbon

import (
	"github.com/gongzhaohui/gef/internal/frontend/components/ribbon/types"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TabContent 组件
type TabContent struct {
	app.Compo
	Tabs          []types.Tab
	ActiveTab     string
	LayoutMode    string
	OnButtonClick func(string)
}

func (t *TabContent) Render() app.UI {
	// 查找当前活动标签
	var activeTab types.Tab
	for _, tab := range t.Tabs {
		if tab.ID == t.ActiveTab {
			activeTab = tab
			break
		}
	}

	// 根据布局模式选择不同的渲染方式
	if t.LayoutMode == "horizontal" {
		return t.renderHorizontal(activeTab)
	}
	return t.renderVertical(activeTab)
}

// 垂直布局渲染
func (t *TabContent) renderVertical(tab types.Tab) app.UI {
	groups := make([]app.UI, len(tab.Groups))
	for i, group := range tab.Groups {
		groups[i] = &RibbonGroup{
			Group:         group,
			LayoutMode:    "vertical",
			OnButtonClick: t.OnButtonClick,
		}
	}

	return app.Div().Class("tab-content").Body(
		app.Div().Class("ribbon-row").Body(groups...),
	)
}

// 水平布局渲染
func (t *TabContent) renderHorizontal(tab types.Tab) app.UI {
	groups := make([]app.UI, len(tab.Groups))
	for i, group := range tab.Groups {
		groups[i] = &RibbonGroup{
			Group:         group,
			LayoutMode:    "horizontal",
			OnButtonClick: t.OnButtonClick,
		}
	}

	return app.Div().Class("tab-content").Body(
		app.Div().Class("ribbon-column").Body(groups...),
	)
}

// RibbonGroup 组件
type RibbonGroup struct {
	app.Compo
	Group         types.Group
	LayoutMode    string
	OnButtonClick func(string)
}

func (r *RibbonGroup) OnMount(ctx app.Context) {
	ctx.Handle("toggleLayout", r.ToggleLayout)
}

// ToggleLayout toggles the layout mode between "horizontal" and "vertical".
func (r *RibbonGroup) ToggleLayout(ctx app.Context, a app.Action) {
	if r.LayoutMode == "horizontal" {
		r.LayoutMode = "vertical"
	} else {
		r.LayoutMode = "horizontal"
	}
	ctx.Update()
}
func (r *RibbonGroup) Render() app.UI {
	className := "ribbon-group"
	itemClassName := "group-item"
	if r.LayoutMode == "horizontal" {
		className += " horizontal"
		itemClassName += " horizontal"
	}

	// 渲染普通按钮
	buttons := make([]app.UI, len(r.Group.Buttons))
	for i, button := range r.Group.Buttons {
		buttons[i] = &RibbonButton{
			Button: button,
			OnClick: func() {
				if r.OnButtonClick != nil {
					r.OnButtonClick(button.ID)
				}
			},
		}
	}

	// 渲染子组
	for _, subGroup := range r.Group.Groups {
		subGroupButtons := make([]app.UI, len(subGroup.Buttons))
		for i, button := range subGroup.Buttons {
			subGroupButtons[i] = &RibbonButton{
				Button: button,
				OnClick: func() {
					if r.OnButtonClick != nil {
						r.OnButtonClick(button.ID)
					}
				},
			}
		}

		buttons = append(buttons,
			app.Div().Class("sub-group").Body(
				app.Div().Class("sub-group-title").Text(subGroup.Name),
				app.Div().Class("sub-group-buttons").Body(subGroupButtons...),
			),
		)
	}

	return app.Div().Class(className).Body(
		app.Div().Class("group-title").Text(r.Group.Title),
		app.Div().Class(itemClassName).Body(buttons...),
	)
}

// RibbonButton 组件
type RibbonButton struct {
	app.Compo
	Button  types.Button
	OnClick func()
}

func (r *RibbonButton) Render() app.UI {
	return app.Button().Class("ribbon-button").Title(r.Button.Title).Body(
		app.I().Class("fa", r.Button.Icon),
		app.Br(),
		app.Span().Text(r.Button.Title),
	).OnClick(func(ctx app.Context, e app.Event) {
		if r.OnClick != nil {
			r.OnClick()
		}
	})
}
