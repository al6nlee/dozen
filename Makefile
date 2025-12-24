# 默认任务
.PHONY: help
help:
	@echo "可用命令："
	@echo "  make dev          - 启动完整开发环境（数据库+应用）"
	@echo "  make db-up        - 启动数据库服务（PostgreSQL + Redis）"
	@echo "  make db-down      - 停止数据库服务"
	@echo "  make app-up       - 启动应用服务（Backend + Frontend）"
	@echo "  make app-down     - 停止应用服务"
	@echo "  make backend      - 本地启动后端服务"
	@echo "  make frontend     - 本地启动前端服务"
	@echo "  make build        - 构建生产版本"
	@echo "  make logs         - 查看所有服务日志"
	@echo "  make db-logs      - 查看数据库服务日志"
	@echo "  make app-logs     - 查看应用服务日志"
	@echo "  make down         - 停止所有服务"
	@echo "  make clean        - 停止所有服务并清理数据"

# 启动数据库服务
.PHONY: db-up
db-up:
	docker compose -f deploy/docker-compose.db.yml up -d
	@echo "数据库服务已启动（PostgreSQL 18 + Redis 8）"

# 停止数据库服务
.PHONY: db-down
db-down:
	docker compose -f deploy/docker-compose.db.yml down
	@echo "数据库服务已停止"

# 启动应用服务
.PHONY: app-up
app-up:
	docker compose -f deploy/docker-compose.yml up -d
	@echo "应用服务已启动"

# 停止应用服务
.PHONY: app-down
app-down:
	docker compose -f deploy/docker-compose.yml down
	@echo "应用服务已停止"

# 启动完整开发环境
.PHONY: dev
dev: db-up
	@echo "等待数据库服务就绪..."
	@sleep 5
	docker compose -f deploy/docker-compose.yml up -d
	@echo "完整开发环境已启动"

# 单独启动后端
.PHONY: backend
backend:
	cd backend && go mod tidy && go run cmd/api/main.go

# 单独启动前端
.PHONY: frontend
frontend:
	cd frontend && npm install && npm run dev

# 构建 Docker 镜像
.PHONY: docker-build
docker-build:
	docker build -t dozen-backend:latest ./backend
	docker build -t dozen-frontend:latest ./frontend

# 查看所有服务日志
.PHONY: logs
logs:
	docker compose -f deploy/docker-compose.db.yml -f deploy/docker-compose.yml logs -f

# 查看数据库服务日志
.PHONY: db-logs
db-logs:
	docker compose -f deploy/docker-compose.db.yml logs -f

# 查看应用服务日志
.PHONY: app-logs
app-logs:
	docker compose -f deploy/docker-compose.yml logs -f

# 停止所有服务
.PHONY: down
down:
	docker compose -f deploy/docker-compose.yml down
	docker compose -f deploy/docker-compose.db.yml down
	@echo "所有服务已停止"

# 清理所有服务和数据
.PHONY: clean
clean:
	docker compose -f deploy/docker-compose.yml down -v
	docker compose -f deploy/docker-compose.db.yml down -v
	rm -rf backend/bin
	rm -rf frontend/dist
	rm -rf frontend/node_modules
	@echo "所有服务已清理"

# 重启所有服务
.PHONY: restart
restart: down dev

# 重启数据库服务
.PHONY: db-restart
db-restart: db-down db-up

# 重启应用服务
.PHONY: app-restart
app-restart: app-down app-up