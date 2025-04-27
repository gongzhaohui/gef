# 项目变量
PROJECT_NAME := gef
FRONTEND_PORT := 8080    # go-app 前端端口
BACKEND_PORT := 3000     # Echo 后端端口

# 开发模式（同时启动前后端）
.PHONY: dev
dev: 
	@echo "🚀 Starting development environment..."
	@echo "Frontend: http://localhost:${FRONTEND_PORT}"
	@echo "API: http://localhost:${BACKEND_PORT}"
	@make -j 2 dev-frontend dev-backend

# 前端开发服务
.PHONY: dev-frontend
dev-frontend:
	@cd cmd/frontend && \
	GOOS=js GOARCH=wasm go build -o ../../web/app.wasm && \
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ../../web/ && \
	go run main.go

# 后端开发服务
.PHONY: dev-backend
dev-backend:
	@cd cmd/backend && \
	go run main.go

# 生产构建
.PHONY: build
build:
	@echo "🔨 Building production binaries..."
	@mkdir -p bin
	@cd cmd/frontend && GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ../../web/app.wasm
	@cd cmd/backend && go build -ldflags "-s -w" -o ../../bin/${PROJECT_NAME}-server
	@cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" web/

# 清理
.PHONY: clean
clean:
	@rm -rf web/app.wasm web/wasm_exec.js bin/

# 帮助
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  dev     - Start frontend (${FRONTEND_PORT}) and backend (${BACKEND_PORT})"
	@echo "  build   - Create production binaries"
	@echo "  clean   - Remove build artifacts"

.DEFAULT_GOAL := help