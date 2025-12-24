# 默认任务
.PHONY: help
help:
	@echo "可用命令："
	@echo "  make dev          - 启动开发环境"
	@echo "  make backend      - 启动后端服务"
	@echo "  make frontend     - 启动前端服务"
	@echo "  make test         - 运行测试"
	@echo "  make build        - 构建生产版本"

# 启动完整开发环境
.PHONY: dev
dev:
	docker-compose -f deploy/docker-compose.yml up -d
	@echo "开发环境已启动"

# 单独启动后端
.PHONY: backend
backend:
	cd backend && go mod tidy && go run cmd/api/main.go

# 单独启动前端
.PHONY: frontend
frontend:
	cd frontend && npm install && npm run dev

# 运行测试
.PHONY: test
test:
	cd backend && go test ./...
	cd frontend && npm test

# 构建 Docker 镜像
.PHONY: docker-build
docker-build:
	docker build -t dozen-backend:latest ./backend
	docker build -t dozen-frontend:latest ./frontend

# 停止开发环境
.PHONY: down
down:
	docker-compose -f deploy/docker-compose.yml down

# 清理
.PHONY: clean
clean:
	docker-compose -f deploy/docker-compose.yml down -v
	rm -rf backend/bin
	rm -rf frontend/dist
	rm -rf frontend/node_modules