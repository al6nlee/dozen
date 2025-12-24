#!/bin/bash

echo "构建 DoZen 项目..."

# 构建后端
echo "构建后端..."
cd backend
go mod tidy
go build -o bin/main cmd/api/main.go
echo "✓ 后端构建完成"

# 构建前端
echo "构建前端..."
cd ../frontend
npm install
npm run build
echo "✓ 前端构建完成"

cd ..
echo "✓ 项目构建完成"