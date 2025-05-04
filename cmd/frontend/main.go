package main

import (
	"encoding/json"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// RibbonMenu 定义Ribbon菜单的数据结构
type RibbonMenu struct {
	Tabs []Tab `json:"tabs"`
}

// Tab 定义Ribbon标签的数据结构
type Tab struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Groups []Group `json:"groups"`
}

// Group 定义Ribbon组的数据结构
type Group struct {
	Title   string   `json:"title"`
	Buttons []Button `json:"buttons"`
}

// Button 定义Ribbon按钮的数据结构
type Button struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// 应用主组件
type OfficeApp struct {
	app.Compo
	activeTab    string
	document     string
	ribbonLayout string
	isDragging   bool
	dragStartX   int
	dragStartY   int
	ribbonMenu   RibbonMenu
	isLoading    bool
	errorMessage string
}

// 初始化
func (a *OfficeApp) OnMount(ctx app.Context) {
	a.activeTab = "home"
	a.document = "New Document"
	a.ribbonLayout = "horizontal"
	a.isLoading = true

	// 异步加载Ribbon菜单数据
	ctx.Async(func() {
		resp, err := http.Get("/web/ribbon_data.json")
		if err != nil {
			a.errorMessage = "Failed to load ribbon data: " + err.Error()
			// ctx.Rebuild()
			return
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&a.ribbonMenu); err != nil {
			a.errorMessage = "Failed to parse ribbon data: " + err.Error()
			// ctx.Rebuild()
			return
		}

		a.isLoading = false
		ctx.Dispatch(func(ctx app.Context) {})
	})
}

// 渲染应用
func (a *OfficeApp) Render() app.UI {
	return app.Div().Class("office-app").Body(
		a.renderTitleBar(),
		a.renderRibbon(),
		a.renderWorkspace(),
	)
}

// 渲染标题栏
func (a *OfficeApp) renderTitleBar() app.UI {
	return app.Header().Class("title-bar").Body(
		app.Div().Class("app-title").Text("Go-App Office"),
		app.Div().Class("document-title").Text(a.document),
		app.Div().Class("window-controls").Body(
			app.Button().Class("window-button minimize").Text("-"),
			app.Button().Class("window-button maximize").Text("□"),
			app.Button().Class("window-button close").Text("×"),
		),
	)
}

// 渲染Ribbon菜单
func (a *OfficeApp) renderRibbon() app.UI {
	if a.isLoading {
		return app.Div().Class("ribbon").Body(
			app.Div().Class("loading-indicator").Text("Loading Ribbon..."),
		)
	}

	if a.errorMessage != "" {
		return app.Div().Class("ribbon").Body(
			app.Div().Class("error-message").Text(a.errorMessage),
		)
	}

	ribbonClass := "ribbon"
	if a.ribbonLayout == "vertical" {
		ribbonClass += " vertical"
	}

	return app.Div().Class(ribbonClass).Body(
		a.renderFileMenu(),
		a.renderTabBar(),
		a.renderDragHandle(),
		a.renderTabContents(),
		a.renderRibbonFooter(),
	)
}

// 渲染拖拽手柄
func (a *OfficeApp) renderDragHandle() app.UI {
	handleClass := "drag-handle"
	if a.ribbonLayout == "vertical" {
		handleClass += " vertical"
	}

	return app.Div().Class(handleClass).
		OnMouseDown(func(ctx app.Context, e app.Event) {
			a.isDragging = true
			a.dragStartX = int(e.Get("clientX").Int())
			a.dragStartY = int(e.Get("clientY").Int())
			e.JSValue().Call("setPointerCapture", int(e.Get("pointerId").Int()))
		}).
		OnMouseMove(func(ctx app.Context, e app.Event) {
			if !a.isDragging {
				return
			}

			currentX := int(e.Get("clientX").Int())
			currentY := int(e.Get("clientY").Int())
			deltaX := currentX - a.dragStartX
			deltaY := currentY - a.dragStartY

			// 根据拖拽方向切换布局
			if abs(deltaX) > abs(deltaY) && a.ribbonLayout == "vertical" {
				a.ribbonLayout = "horizontal"
				a.isDragging = false
				// ctx.Rebuild()
			} else if abs(deltaY) > abs(deltaX) && a.ribbonLayout == "horizontal" {
				a.ribbonLayout = "vertical"
				a.isDragging = false
				// ctx.Rebuild()
			}
		}).
		OnMouseUp(func(ctx app.Context, e app.Event) {
			a.isDragging = false
			e.JSValue().Call("releasePointerCapture", int(e.Get("pointerId").Int()))
		})
}

// 渲染文件菜单
func (a *OfficeApp) renderFileMenu() app.UI {
	return app.Div().Class("file-menu").Body(
		app.Button().Class("file-button").Text("File").
			OnClick(func(ctx app.Context, e app.Event) {
				// 实现文件菜单逻辑
				if a.ribbonLayout == "vertical" {
					a.ribbonLayout = "horizontal"
				} else {
					a.ribbonLayout = "vertical"
				}
			}),
	)
}

// 渲染标签栏
func (a *OfficeApp) renderTabBar() app.UI {
	tabBarClass := "tab-bar"
	if a.ribbonLayout == "vertical" {
		tabBarClass += " vertical"
	}

	tabs := make([]app.UI, len(a.ribbonMenu.Tabs))
	for i, tab := range a.ribbonMenu.Tabs {
		tabs[i] = a.renderDynamicTab(tab)
	}

	return app.Div().Class(tabBarClass).Body(tabs...)
}

// 渲染动态标签
func (a *OfficeApp) renderDynamicTab(tab Tab) app.UI {
	className := "tab"
	if a.activeTab == tab.ID {
		className += " active"
	}

	return app.Button().Class(className).ID(tab.ID).Text(tab.Title).
		OnClick(func(ctx app.Context, e app.Event) {
			a.activeTab = tab.ID
			// ctx.Rebuild()
		})
}

// 渲染标签内容
func (a *OfficeApp) renderTabContents() app.UI {
	// 查找当前活动标签
	for _, tab := range a.ribbonMenu.Tabs {
		if tab.ID == a.activeTab {
			return a.renderTabContent(tab)
		}
	}

	return app.Div().Text("Tab not found")
}

// 渲染标签内容
func (a *OfficeApp) renderTabContent(tab Tab) app.UI {
	layoutClass := "ribbon-row"
	if a.ribbonLayout == "vertical" {
		layoutClass = "ribbon-column"
	}

	groups := make([]app.UI, len(tab.Groups))
	for i, group := range tab.Groups {
		groups[i] = a.renderDynamicGroup(group)
	}

	return app.Div().Class(layoutClass).Body(groups...)
}

// 渲染动态Ribbon组
func (a *OfficeApp) renderDynamicGroup(group Group) app.UI {
	groupClass := "ribbon-group"
	if a.ribbonLayout == "vertical" {
		groupClass += " vertical"
	}

	buttons := make([]app.UI, len(group.Buttons))
	for i, button := range group.Buttons {
		buttons[i] = a.renderDynamicButton(button)
	}

	return app.Div().Class(groupClass).Body(
		app.Div().Class("group-title").Text(group.Title),
		app.Div().Class("group-items").Body(buttons...),
	)
}

// 渲染动态Ribbon按钮
func (a *OfficeApp) renderDynamicButton(button Button) app.UI {
	buttonClass := "ribbon-button"
	if a.ribbonLayout == "vertical" {
		buttonClass += " vertical"
	}

	return app.Button().Class(buttonClass).ID(button.ID).Body(
		app.I().Class("fa", button.Icon),
		app.Br(),
		app.Span().Text(button.Title),
	).OnClick(func(ctx app.Context, e app.Event) {
		a.handleRibbonAction(button.ID)
	})
}

// 处理Ribbon按钮动作
func (a *OfficeApp) handleRibbonAction(id string) {
	// 根据按钮ID执行相应操作
	app.Log("Ribbon action:", id)

	// 示例：处理粘贴操作
	if id == "paste" {
		// 实现粘贴逻辑
	}
}

// 渲染Ribbon底部
func (a *OfficeApp) renderRibbonFooter() app.UI {
	return app.Div().Class("ribbon-footer").Body(
		app.Button().Class("minimize-button").Text("^").
			OnClick(func(ctx app.Context, e app.Event) {
				// 实现最小化Ribbon逻辑
			}),
	)
}

// 渲染工作区
func (a *OfficeApp) renderWorkspace() app.UI {
	return app.Div().Class("workspace").Body(
		app.Textarea().Class("document-editor").Text(a.document).
			OnChange(func(ctx app.Context, e app.Event) {
				a.document = ctx.JSSrc().Get("value").String()
			}),
	)
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	app.Route("/", app.NewZeroComponentFactory(&OfficeApp{}))
	http.Handle("/", &app.Handler{
		Name:        "OfficeApp",
		Title:       "OfficeApp",
		Description: "An OfficeApp example",
	})

	// app.Handle("/web/ribbon_data.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	data, err := os.ReadFile("ribbon_data.json")
	// 	if err != nil {
	// 		http.Error(w, "Failed to read ribbon data", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(data)
	// })
	app.RunWhenOnBrowser()
}
