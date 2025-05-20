# OfficeApp

OfficeApp 是一个基于 **全栈 Go** 和 **声明式 UI** 的示例项目，使用 [go-app](https://github.com/maxence-charriere/go-app) 框架构建。它展示了如何使用 Go 编写前端和后端代码，构建现代 Web 应用程序。

## 项目结构

```folder
gef/
├── cmd/                  # 主入口[5,9](@ref)
│   ├── server/           # 后端主程序
│   │   └── main.go       # Echo服务器启动入口
│   └── webapp/           # 前端编译入口（如需服务端渲染）
│       └── main.go       # go-app编译入口
├── internal/             # 私有代码隔离[1,5](@ref)
│   ├── api/              # API定义（Protobuf/OpenAPI）
│   ├── app/              # 前后端共享逻辑
│   │   ├── components/   # 通用组件（DTO/验证器等）[9](@ref)
│   │   └── utils/        # 跨层工具函数
│   ├── backend/          # 后端核心模块[9,10](@ref)
│   │   ├── config/       # 配置加载
│   │   ├── controllers/   # Echo路由处理器
│   │   ├── middlewares/   # 中间件（JWT/日志等）
│   │   ├── services/      # 业务逻辑层
│   │   ├── repositories/   # GORM数据访问层
│   │   └── models/        # 数据库模型
│   └── frontend/         # 前端核心模块
│       ├── components/   # go-app组件
│       ├── routes/       # 前端路由定义
│       └── store/        # 状态管理（如Redux模式）
├── web/                  # 前端资源[1,5](@ref)
│   ├── static/           # 静态资源（CSS/图片）
│   ├── wasm/             # go-app生成的WebAssembly
│   └── templates/        # 服务端模板（可选）
├── migrations/           # 数据库迁移脚本[9](@ref)
├── pkg/                  # 可复用公共库[1](@ref)
├── scripts/              # 构建部署脚本[5](@ref)
├── docs/                 # 项目文档
├── test/                 # 测试套件
│   ├── e2e/              # 端到端测试
│   └── mocks/            # Mock数据
├── go.mod
└── Makefile              # 统一构建命令[5](@ref)
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
