
.PHONY: up down build logs-backend logs-frontend migrate seed test-backend test-frontend shell-backend shell-frontend shell-db clean help

# デフォルトのターゲット
.DEFAULT_GOAL := help

# 変数定義
DOCKER_COMPOSE = docker-compose
BACKEND_SERVICE = backend
FRONTEND_SERVICE = frontend
DB_SERVICE = postgres

# 全てのサービスを起動
up:
	$(DOCKER_COMPOSE) up -d
	@echo "All services are up and running"
	@echo "Frontend: http://localhost:3000"
	@echo "Backend API: http://localhost:8080/api"

# データベースのみ起動
up-db:
	$(DOCKER_COMPOSE) up -d $(DB_SERVICE)
	@echo "Database is running on port 5432"

# バックエンドのみ起動
up-backend: up-db
	$(DOCKER_COMPOSE) up -d $(BACKEND_SERVICE)
	@echo "Backend is running on http://localhost:8080/api"

# フロントエンドのみ起動
up-frontend:
	$(DOCKER_COMPOSE) up -d $(FRONTEND_SERVICE)
	@echo "Frontend is running on http://localhost:3000"

# サービスを停止
down:
	$(DOCKER_COMPOSE) down
	@echo "All services are stopped"

# サービスをビルド
build:
	$(DOCKER_COMPOSE) build
	@echo "All services are built"

# バックエンドのログを表示
logs-backend:
	$(DOCKER_COMPOSE) logs -f $(BACKEND_SERVICE)

# フロントエンドのログを表示
logs-frontend:
	$(DOCKER_COMPOSE) logs -f $(FRONTEND_SERVICE)

# マイグレーションを実行
migrate:
	$(DOCKER_COMPOSE) exec $(DB_SERVICE) psql -U postgres -d jobhunting_db -f /app/migrations/schema.sql
	@echo "Database migrations completed"

# テストデータを投入
seed:
	$(DOCKER_COMPOSE) exec $(DB_SERVICE) psql -U postgres -d jobhunting_db -f /app/scripts/seed.sql
	@echo "Test data seeded into database"

# バックエンドのテスト実行
test-backend:
	$(DOCKER_COMPOSE) exec $(BACKEND_SERVICE) go test ./...

# フロントエンドのテスト実行
test-frontend:
	$(DOCKER_COMPOSE) exec $(FRONTEND_SERVICE) npm test

# バックエンドのシェルにアクセス
shell-backend:
	$(DOCKER_COMPOSE) exec $(BACKEND_SERVICE) sh

# フロントエンドのシェルにアクセス
shell-frontend:
	$(DOCKER_COMPOSE) exec $(FRONTEND_SERVICE) sh

# データベースのシェルにアクセス
shell-db:
	$(DOCKER_COMPOSE) exec $(DB_SERVICE) psql -U postgres -d jobhunting_db

# 一時ファイルやコンテナを削除
clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans
	rm -rf backend/tmp/* frontend/.next frontend/node_modules

# ヘルプを表示
help:
	@echo "Available commands:"
	@echo "  make up               - Start all services"
	@echo "  make up-db            - Start only database"
	@echo "  make up-backend       - Start database and backend"
	@echo "  make up-frontend      - Start frontend"
	@echo "  make down             - Stop all services"
	@echo "  make build            - Build all services"
	@echo "  make logs-backend     - Show backend logs"
	@echo "  make logs-frontend    - Show frontend logs"
	@echo "  make migrate          - Run database migrations"
	@echo "  make seed             - Seed test data"
	@echo "  make test-backend     - Run backend tests"
	@echo "  make test-frontend    - Run frontend tests"
	@echo "  make shell-backend    - Access backend shell"
	@echo "  make shell-frontend   - Access frontend shell"
	@echo "  make shell-db         - Access database shell"
	@echo "  make clean            - Clean temporary files and containers"
	@echo "  make help             - Show this help message"