package widgets

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app" // 导入go-app包
)

// TitleBar 组件
type TitleBar struct {
	app.Compo
	DocumentTitle string
}

func (t *TitleBar) Render() app.UI {
	return app.Div().Class("title-bar").Body(
		app.Div().Class("app-title").Text("Office 2010 Style App"),
		app.Div().Class("document-title").Text(t.DocumentTitle),
		app.Div().Class("window-controls").Body(
			app.Button().Class("window-button").Title("Minimize").Text("—"),
			app.Button().Class("window-button").Title("Maximize").Text("□"),
			app.Button().Class("window-button close").Title("Close").Text("×"),
		),
	)
}

// TitleBar 组件用于显示应用程序的标题栏，包含应用程序标题、文档标题和窗口控制按钮。
// 它使用了go-app包来创建和管理组件。TitleBar组件的Render方法返回一个包含标题栏内容的UI元素。
