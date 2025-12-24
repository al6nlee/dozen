#!/bin/bash

echo "启动 DoZen 开发环境..."

# 检查 Docker 是否运行
if ! docker info > /dev/null 2>&1; then
  echo "错误: Docker 未运行，请先启动 Docker"
  exit 1
fi

# 启动服务
docker-compose -f deploy/docker-compose.yml up -d

echo "✓ 开发环境已启动"
echo "  - 后端: http://localhost:8080"
echo "  - 前端: http://localhost:3000"
echo "  - 数据库: localhost:5432"