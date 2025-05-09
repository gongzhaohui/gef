package widgets

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// StatusBar 组件
type StatusBar struct {
	app.Compo
	Document string
}

func (s *StatusBar) Render() app.UI {
	return app.Div().Class("status-bar").Body(
		app.Div().Class("status-info").Text("Ready"),
		app.Div().Class("document-info").Text(s.Document),
	)
}
