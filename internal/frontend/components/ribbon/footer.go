package ribbon

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// RibbonFooter 组件
type RibbonFooter struct {
	app.Compo
	IsCollapsed      bool
	OnToggleCollapse func()
	RibbonPosition   string
}

func (r *RibbonFooter) OnMount(ctx app.Context) {
	ctx.Handle("toggleLayout", r.ToggleLayout)
}

// ToggleLayout toggles the collapsed state of the ribbon footer.
func (r *RibbonFooter) ToggleLayout(ctx app.Context, a app.Action) {
	r.IsCollapsed = !r.IsCollapsed
	ctx.Update()
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

	// 下部Ribbon的折叠按钮方向相反
	if r.RibbonPosition == "lower" {
		buttonIcon = "fa-chevron-down"
		if r.IsCollapsed {
			buttonIcon = "fa-chevron-up"
		}
	}

	return app.Div().Class("ribbon-footer").Body(
		app.Button().Class(buttonClass).Title(buttonTitle).Body(
			app.I().Class("fa", buttonIcon),
		).OnClick(func(ctx app.Context, e app.Event) {
			if r.OnToggleCollapse != nil {
				r.OnToggleCollapse()
			}
		}),
	)
}
