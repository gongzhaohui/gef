package types

// RibbonMenu 定义Ribbon菜单的数据结构
type RibbonMenu struct {
	Tabs []Tab `json:"tabs"`
}

// Tab 定义Ribbon标签的数据结构
type Tab struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Groups []Group `json:"groups"`
}

// Group 定义Ribbon组的数据结构
type Group struct {
	Title     string     `json:"title"`
	Name      string     `json:"name"`
	Buttons   []Button   `json:"buttons"`
	SubGroups []SubGroup `json:"subGroups,omitempty"`
}

// SubGroup 定义Ribbon子组的数据结构
type SubGroup struct {
	Name    string   `json:"name"`
	Buttons []Button `json:"buttons"`
}

// Button 定义Ribbon按钮的数据结构
type Button struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Group string `json:"group,omitempty"` // 所属子组
}
