# OfficeApp

OfficeApp 是一个基于 **全栈 Go** 和 **声明式 UI** 的示例项目，使用 [go-app](https://github.com/maxence-charriere/go-app) 框架构建。它展示了如何使用 Go 编写前端和后端代码，构建现代 Web 应用程序。

## 项目结构

```folder
e:\github\gef\
├── cmd/
│   └── frontend/          # 前端主程序
│       └── main.go        # 应用入口
├── pkg/
│   └── components/
│       └── widgets/       # 自定义组件（如 TitleBar、Receptacle、StatusBar 等）
│   └── services/          # 服务层（如数据集服务）
├── web/
│   └── ribbon_data.json   # 前端 UI 配置文件
└── README.md              # 项目说明
```

## 技术栈

- **全栈 Go**：前端和后端均使用 Go 编写，简化了开发流程。
- **go-app**：用于构建声明式 UI 的 Go 框架，支持现代 Web 应用的开发。
- **组件化设计**：通过自定义组件（如 `TitleBar`、`Receptacle`、`StatusBar`）实现模块化开发。
- **JSON 配置**：通过 `ribbon_data.json` 文件动态配置前端 UI。

## 功能特性

1. **声明式 UI**：
   - 使用 `go-app` 框架，通过 Go 代码直接定义前端界面。
   - 示例代码：

    ```go
     return app.Div().Class("office-app", layoutClass).Body(
         &widgets.TitleBar{
             DocumentTitle:  a.document,
             OnLayoutToggle: a.toggleLayout,
         },
         &widgets.Receptacle{
             LayoutMode: a.layoutMode,
         },
         &widgets.StatusBar{
             Document: a.document,
         },
     )
    ```

2. **动态布局切换**：
   - 支持垂直（`vertical`）和水平（`horizontal`）布局模式。
   - 切换逻辑：

    ```go
     func (a *OfficeApp) toggleLayout(ctx app.Context) {
         if a.layoutMode == "vertical" {
             a.layoutMode = "horizontal"
         } else {
             a.layoutMode = "vertical"
         }
         log.Printf("Layout mode changed to: %v", a.layoutMode)
     }
    ```

3. **组件化设计**：
   - 自定义组件如 `TitleBar`、`Receptacle` 和 `StatusBar`，实现代码复用和模块化开发。

4. **JSON 配置驱动**：
   - 使用 `ribbon_data.json` 文件配置前端按钮和操作，支持动态扩展。

5. **全栈开发**：
   - 前端和后端均使用 Go 编写，简化了技术栈，提升了开发效率。

## 快速开始

### 环境要求

- Go 1.20 或更高版本

### 运行项目

1. 克隆项目：

   ```bash
   git clone https://github.com/your-repo/gef.git
   cd gef
   ```

2. 启动应用：

   ```bash
   #go mod tidy
   make dev
   ```

3. 在浏览器中访问：

   ```url
   http://localhost:3000
   ```

## 项目组件

### TitleBar

- 显示文档标题，并提供布局切换按钮。
- 示例：

  ```go
  &widgets.TitleBar{
      DocumentTitle:  a.document,
      OnLayoutToggle: a.toggleLayout,
  }
  ```

### Receptacle

- 主内容区域，支持垂直和水平布局。
- 示例：

  ```go
  &widgets.Receptacle{
      LayoutMode: a.layoutMode,
  }
  ```

### StatusBar

- 显示文档状态信息。
- 示例：

  ```go
  &widgets.StatusBar{
      Document: a.document,
  }
  ```

## 配置文件

ribbon_data.json 文件用于配置前端按钮和操作。例如：

```json
{
    "id": "alignLeft",
    "title": "Align Left",
    "icon": "fa-align-left",
    "action": "alignLeftAction",
    "actionPath": "/paragraph/alignLeft"
}
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进此项目！

## 许可证

此项目使用 MIT License 进行许可。有关更多信息，请参见 [LICENSE](LICENSE) 文件。
