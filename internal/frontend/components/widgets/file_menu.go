package widgets

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// FileMenu represents the file menu component.
type FileMenu struct {
	app.Compo
	OnLayoutToggle func(ctx app.Context)
}

// Render renders the FileMenu component.
func (f *FileMenu) Render() app.UI {
	return app.Div().Class("file-menu").Body(
		app.Button().
			Class("layout-toggle").
			Text("Toggle Layout").
			OnClick(func(ctx app.Context, e app.Event) {
				if f.OnLayoutToggle != nil {
					f.OnLayoutToggle(ctx)
				}
			}),
	)
}
