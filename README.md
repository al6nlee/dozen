## 项目介绍

> 该项目名叫DoZen，意为**实现禅道**。这是一个基于Go和React的Web应用程序。

## 目录结构

> 采用Monorepo（单仓库多项目）的目录结构

```angular2html
dozen/
├── backend/                    # Go 后端
│   ├── cmd/
│   │   └── api/               # 主程序入口
│   │       └── main.go
│   ├── internal/              # 私有应用代码
│   │   ├── handler/           # HTTP handlers
│   │   ├── service/           # 业务逻辑层
│   │   ├── repository/        # 数据访问层
│   │   ├── model/             # 数据模型
│   │   ├── middleware/        # 中间件
│   │   └── config/            # 配置管理
│   ├── pkg/                   # 可复用的公共库
│   │   ├── logger/
│   │   ├── database/
│   │   └── utils/
│   ├── api/                   # API 定义
│   │   └── openapi/           # OpenAPI/Swagger 规范
│   ├── migrations/            # 数据库迁移文件
│   ├── test/                  # 集成测试
│   ├── scripts/               # 构建/部署脚本
│   ├── go.mod
│   ├── go.sum
│   ├── Dockerfile
│   └── .env.example
├── frontend/                  # React 前端
│   ├── public/
│   ├── src/
│   │   ├── components/        # 通用组件
│   │   ├── pages/             # 页面组件
│   │   ├── hooks/             # 自定义 Hooks
│   │   ├── services/          # API 调用封装
│   │   ├── store/             # 状态管理 (Redux/Zustand)
│   │   ├── utils/             # 工具函数
│   │   ├── types/             # TypeScript 类型定义
│   │   ├── assets/            # 静态资源
│   │   ├── App.tsx
│   │   └── main.tsx
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts         # 或 webpack.config.js
│   ├── Dockerfile
│   └── .env.example
├── deploy/                    # 部署配置
│   ├── kubernetes/
│   │   ├── backend/
│   │   │   ├── deployment.yaml
│   │   │   ├── service.yaml
│   │   │   └── ingress.yaml
│   │   └── frontend/
│   │       ├── deployment.yaml
│   │       └── service.yaml
│   └── docker-compose.yml     # 本地开发环境
├── docs/                      # 项目文档
│   ├── api/                   # API 文档
│   └── architecture/          # 架构设计文档
├── scripts/                   # 全局脚本
│   ├── dev.sh                 # 本地开发启动
│   └── build.sh               # 构建脚本
├── .gitignore
├── Makefile                   # 统一的构建命令
└── README.md
```
