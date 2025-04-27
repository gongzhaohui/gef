package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Pagination struct {
	app.Compo
	CurrentPage int
	PageSize    int
	Total       int
}
