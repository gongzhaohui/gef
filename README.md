# GEF

GEF is a full-stack Go application that demonstrates the use of [go-app](https://github.com/maxence-charriere/go-app) for building declarative UI components. This project showcases how to create modern web applications using Go for both frontend and backend development.

## Project Structure

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
