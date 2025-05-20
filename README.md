# GEF

GEF is a full-stack Go application that demonstrates the use of [go-app](https://github.com/maxence-charriere/go-app) for building declarative UI components. This project showcases how to create modern web applications using Go for both frontend and backend development.

## Project Structure

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
│   │   └── model/        # 数据库模型
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
├── tests/                 # 测试套件
│   ├── e2e/              # 端到端测试
│   └── mocks/            # Mock数据
├── go.mod
└── Makefile              # 统一构建命令[5](@ref)
```

## Features

1. **Declarative UI**:
   - Built with `go-app`, enabling declarative UI development directly in Go.
   - Example:

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

2. **Dynamic Layout Switching**:
   - Supports vertical (`vertical`) and horizontal (`horizontal`) layout modes.
   - Example toggle logic:

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

3. **Component-Based Design**:
   - Modular components like `TitleBar`, `Receptacle`, and `StatusBar` for reusable and maintainable code.

4. **JSON-Driven Configuration**:
   - UI elements and actions are dynamically configured using `ribbon_data.json`.

5. **Full-Stack Go**:
   - Both frontend and backend are implemented in Go, simplifying the development stack.

## Quick Start

### Prerequisites

- Go 1.20 or later

### Run the Application

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/gef.git
   cd gef
   ```

2. Start the application:

    ```bash
    #go mod tidy
    make dev
    ```

3. Open your browser and navigate to:

    ```url
    http://localhost:3000
    ```

## Project Components

### TitleBar

- Displays the document title and provides a layout toggle button.
- Example:

    ```go
    &widgets.TitleBar{
        DocumentTitle:  a.document,
        OnLayoutToggle: a.toggleLayout,
    }
    ```

### Receptacle

- Main content area supporting vertical and horizontal layouts.
- Example:

    ```go
    &widgets.Receptacle{
        LayoutMode: a.layoutMode,
    }
    ```

### StatusBar

- Displays document status information.
- Example:

    ```go
    &widgets.StatusBar{
        Document: a.document,
    }
    ```

## Configuration

The web/ribbon_data.json file is used to configure frontend buttons and actions. Example:

```json
{
    "buttons": [
        {
            "id": "button1",
            "label": "Button 1",
            "icon": "fas fa-plus",
            "action": "action1"
        },
        {
            "id": "button2",
            "label": "Button 2",
            "icon": "fas fa-minus",
            "action": "action2"
        }
    ]
}
```

Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
