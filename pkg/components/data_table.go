package components

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type DataTable struct {
	app.Compo
	data    []Order
	Loading bool   // 加载状态
	Error   string // 错误信息
}

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}

func (c *DataTable) OnMount(ctx app.Context) {
	c.loadData(ctx)
}

func (c *DataTable) loadData(ctx app.Context) {
	c.Loading = true

	c.Error = ""

	ctx.Async(func() { // Use framework's async handling
		req, _ := http.NewRequest("GET", "/api/orders", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.Dispatch(func(ctx app.Context) {
				c.Error = fmt.Sprintf("Request failed: %v", err)
				c.Loading = false
			})
			return
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			ctx.Dispatch(func(ctx app.Context) {
				c.Error = fmt.Sprintf("API error (%d): %s", resp.StatusCode, string(data))
				c.Loading = false
			})
			return
		}

		var orders []Order
		if err := json.Unmarshal(data, &orders); err != nil {
			ctx.Dispatch(func(ctx app.Context) {
				c.Error = fmt.Sprintf("Data parsing failed: %v", err)
				c.Loading = false
			})
			return
		}

		ctx.Dispatch(func(ctx app.Context) {
			c.data = orders
			c.Loading = false
		})
	})

}

func (c *DataTable) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("订单数据表"),
		c.renderLoader(),
		c.renderTable(),
	)
}

func (c *DataTable) renderLoader() app.UI {
	if !c.Loading {
		return nil
	}
	return app.Div().
		Class("loader").
		Text("加载中...")
}

func (c *DataTable) renderTable() app.UI {
	if len(c.data) == 0 {
		return app.P().Text("暂无数据")
	}

	rows := make([]app.UI, len(c.data))
	for i, order := range c.data {
		rows[i] = app.Tr().Body(
			app.Td().Text(order.ID),
			app.Td().Text(order.Amount),
			app.Td().Text(order.Status),
		)
	}

	return app.Table().Body(
		app.THead().Body(
			app.Tr().Body(
				app.Th().Text("订单ID"),
				app.Th().Text("金额"),
				app.Th().Text("状态"),
			),
		),
		app.TBody().Body(rows...),
	)
}
