package nested_table

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Column defines the structure of a table column.
type Column struct {
	Title    string
	Key      string
	Sortable bool                                     // Indicates if the column is sortable
	Renderer func(data map[string]interface{}) string // Custom renderer for cell content
}

// NestedRow represents a row with potential nested rows.
type NestedRow struct {
	Data     map[string]interface{}
	Children []*NestedRow
}

// NestedTable represents a table with nested rows and column definitions.
type NestedTable struct {
	app.Compo
	Columns      []Column
	Rows         []*NestedRow
	SortedColumn string
	SortAsc      bool
	PageSize     int
	CurrentPage  int
}

// NewNestedTable creates a new NestedTable with the specified columns.
func NewNestedTable(columns []Column, pageSize int) *NestedTable {
	return &NestedTable{
		Columns:     columns,
		Rows:        make([]*NestedRow, 0),
		PageSize:    pageSize,
		CurrentPage: 1,
	}
}

// AddRow adds a new row to the NestedTable.
func (nt *NestedTable) AddRow(row *NestedRow) {
	nt.Rows = append(nt.Rows, row)
}

// SortRows sorts the rows based on the specified column key.
func (nt *NestedTable) SortRows(columnKey string) {
	if nt.SortedColumn == columnKey {
		nt.SortAsc = !nt.SortAsc
	} else {
		nt.SortedColumn = columnKey
		nt.SortAsc = true
	}

	sort.Slice(nt.Rows, func(i, j int) bool {
		val1, ok1 := nt.Rows[i].Data[columnKey]
		val2, ok2 := nt.Rows[j].Data[columnKey]
		if !ok1 || !ok2 {
			return false
		}
		if nt.SortAsc {
			return fmt.Sprintf("%v", val1) < fmt.Sprintf("%v", val2)
		}
		return fmt.Sprintf("%v", val1) > fmt.Sprintf("%v", val2)
	})
}

// PaginateRows returns the rows for the current page.
func (nt *NestedTable) PaginateRows() []*NestedRow {
	start := (nt.CurrentPage - 1) * nt.PageSize
	end := start + nt.PageSize
	if start >= len(nt.Rows) {
		return []*NestedRow{}
	}
	if end > len(nt.Rows) {
		end = len(nt.Rows)
	}
	return nt.Rows[start:end]
}

// renderRow recursively renders a row and its nested rows.
func (nt *NestedTable) renderRow(sb *strings.Builder, row *NestedRow, level int) {
	sb.WriteString("<tr>\n")
	for _, col := range nt.Columns {
		content := ""
		if col.Renderer != nil {
			content = col.Renderer(row.Data)
		} else if value, ok := row.Data[col.Key]; ok {
			content = fmt.Sprintf("%v", value)
		}
		sb.WriteString(fmt.Sprintf("<td style='padding-left:%dpx'>%s</td>\n", level*20, content))
	}
	sb.WriteString("</tr>\n")

	// Render nested rows
	for _, child := range row.Children {
		nt.renderRow(sb, child, level+1)
	}
}

// Render generates the HTML for the NestedTable.
func (nt *NestedTable) Render() app.UI {
	var sb strings.Builder
	sb.WriteString("<table class='ant-table'>\n")
	sb.WriteString("<thead class='ant-table-thead'>\n<tr>\n")
	for _, col := range nt.Columns {
		if col.Sortable {
			sb.WriteString(fmt.Sprintf(
				"<th onclick='sort(\"%s\")'>%s <span>%s</span></th>\n",
				col.Key,
				col.Title,
				nt.getSortIndicator(col.Key),
			))
		} else {
			sb.WriteString(fmt.Sprintf("<th>%s</th>\n", col.Title))
		}
	}
	sb.WriteString("</tr>\n</thead>\n<tbody class='ant-table-tbody'>\n")
	for _, row := range nt.PaginateRows() {
		nt.renderRow(&sb, row, 0)
	}
	sb.WriteString("</tbody>\n</table>\n")

	// Pagination controls
	sb.WriteString(nt.renderPagination())

	return app.Raw(sb.String())
}

// getSortIndicator returns the sort indicator for a column.
func (nt *NestedTable) getSortIndicator(columnKey string) string {
	if nt.SortedColumn == columnKey {
		if nt.SortAsc {
			return "▲"
		}
		return "▼"
	}
	return ""
}

// renderPagination renders the pagination controls.
func (nt *NestedTable) renderPagination() string {
	var sb strings.Builder
	totalPages := (len(nt.Rows) + nt.PageSize - 1) / nt.PageSize
	sb.WriteString("<div class='ant-pagination'>\n")
	for i := 1; i <= totalPages; i++ {
		if i == nt.CurrentPage {
			sb.WriteString(fmt.Sprintf("<span class='ant-pagination-item ant-pagination-item-active'>%d</span>\n", i))
		} else {
			sb.WriteString(fmt.Sprintf("<span class='ant-pagination-item' onclick='goToPage(%d)'>%d</span>\n", i, i))
		}
	}
	sb.WriteString("</div>\n")
	return sb.String()
}
