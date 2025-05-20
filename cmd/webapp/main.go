package main

import (
	"log"
	"net/http"

	"github.com/gongzhaohui/gef/internal/frontend/components/widgets"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// 应用主组件
type OfficeApp struct {
	app.Compo
	// activeTab              string
	document   string
	layoutMode string // "vertical" 或 "horizontal"
}

// 初始化
func (a *OfficeApp) OnMount(ctx app.Context) {
	a.layoutMode = "vertical" // 默认垂直布局（上中下）
	a.document = "New Document"
	// ctx.ObserveState("layoutMode", &a.layoutMode)

}

// 渲染应用
func (a *OfficeApp) Render() app.UI {
	// 根据布局模式应用不同的类
	layoutClass := "app-layout-vertical"

	return app.Div().Class("office-app", layoutClass).Body(
		&widgets.TitleBar{
			DocumentTitle:  a.document,
			OnLayoutToggle: a.toggleLayout,
		},
		// &widgets.FileMenu{
		// 	OnLayoutToggle: a.toggleLayout,
		// },
		&widgets.Receptacle{
			LayoutMode: a.layoutMode,
		},
		// 状态栏
		&widgets.StatusBar{
			Document: a.document,
		},
	)
}
func (a *OfficeApp) toggleLayout(ctx app.Context) {

	// 切换布局模式
	log.Printf("app.toggle Layout mode : %v", a.layoutMode)
	if a.layoutMode == "vertical" {
		a.layoutMode = "horizontal"
	} else {
		a.layoutMode = "vertical"
	}
	// ctx.SetState("layoutMode", a.layoutMode)
	// ctx.Update()
	log.Printf("Layout mode changed to: %v", a.layoutMode)
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
