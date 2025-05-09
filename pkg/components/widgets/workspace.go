package widgets

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Workspace 组件
type Workspace struct {
	app.Compo
	Document string
	OnChange func(string)
}

func (w *Workspace) Render() app.UI {
	return app.Div().Class("workspace").Body(
		app.Textarea().Class("document-editor").Text(w.Document).
			OnChange(func(ctx app.Context, e app.Event) {
				if w.OnChange != nil {
					w.OnChange(ctx.JSSrc().Get("value").String())
				}
			}),
	)
}
