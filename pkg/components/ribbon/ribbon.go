package ribbon

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type IItem struct {
	Caption string
	Icon    string
	Link    string
	Group   *IGroup
}
type IGroup struct {
	Caption string
	Items   []*IItem
}
type Ribbon struct {
	app.Compo
	MenuItems []MenuItem

	Error   string
	Loading bool
}

func (c *Ribbon) OnMount(ctx app.Context) {

	if c.Loading {
		return
	}
	c.Loading = true
	defer func() {
		c.Loading = false
	}()
	ctx.Async(func() {
		getMenuData()

	})
}

func getMenuData() {
	panic("unimplemented")
}
