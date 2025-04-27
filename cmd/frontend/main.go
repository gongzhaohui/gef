package main

import (
	"gef/pkg/components/nested_table"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type MyApp struct {
	app.Compo
}

func (m *MyApp) Render() app.UI {
	columns := []nested_table.Column{
		{Title: "Name", Key: "name", Sortable: true},
		{Title: "Age", Key: "age", Sortable: true},
		{Title: "Address", Key: "address"},
	}

	table := nested_table.NewNestedTable(columns, 5)

	// 添加数据
	table.AddRow(&nested_table.NestedRow{
		Data: map[string]interface{}{"name": "John", "age": 30, "address": "123 Main St"},
	})
	table.AddRow(&nested_table.NestedRow{
		Data: map[string]interface{}{"name": "Jane", "age": 28, "address": "456 Elm St"},
	})

	return table.Render()
}

func main() {
	app.Route("/", func() app.Composer { return &MyApp{} })
	app.RunWhenOnBrowser()
}
