.PHONY: help setup dev stop restart logs clean db-shell test

# Colors
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
BLUE := \033[0;34m
NC := \033[0m

help: ## ヘルプを表示
	@echo "$(BLUE)LabCareer Development Commands:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(YELLOW)Quick Start:$(NC)"
	@echo "  1. make setup    # 初回セットアップ"
	@echo "  2. make dev      # 開発サーバー起動"

setup: ## 初回セットアップ
	@echo "$(YELLOW)📦 Setting up LabCareer development environment...$(NC)"
	@cp -n .env.example .env 2>/dev/null || true
	@echo "$(GREEN)✓ Environment file created$(NC)"
	@cd frontend && npm install
	@echo "$(GREEN)✓ Frontend dependencies installed$(NC)"
	@cd backend && go mod download
	@echo "$(GREEN)✓ Backend dependencies downloaded$(NC)"
	@docker-compose build
	@echo "$(GREEN)✓ Docker images built$(NC)"
	@echo ""
	@echo "$(GREEN)✅ Setup complete! Run 'make dev' to start development servers$(NC)"

dev: ## 開発サーバーを起動（Docker: Backend+DB, Local: Frontend）
	@echo "$(YELLOW)🚀 Starting development servers...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)✓ Backend and Database started (Docker)$(NC)"
	@echo ""
	@echo "$(BLUE)Next step:$(NC)"
	@echo "  Open a new terminal and run:"
	@echo "  $(GREEN)cd frontend && npm run dev$(NC)"
	@echo ""
	@echo "$(YELLOW)Services:$(NC)"
	@echo "  Frontend: http://localhost:3000 (run manually)"
	@echo "  Backend:  http://localhost:8080 (Docker)"
	@echo "  Database: localhost:5432 (Docker)"

dev-all: ## 全サービスを起動（別ターミナルでフロントエンドも起動）
	@make dev
	@echo "$(YELLOW)Starting frontend in 3 seconds...$(NC)"
	@sleep 3
	@cd frontend && npm run dev

backend: ## バックエンドとDBのみ起動
	@docker-compose up

backend-d: ## バックエンドとDBをバックグラウンドで起動
	@docker-compose up -d
	@echo "$(GREEN)✓ Backend services started$(NC)"

frontend: ## フロントエンドのみ起動（ローカル）
	@cd frontend && npm run dev

stop: ## 全サービスを停止
	@docker-compose stop
	@echo "$(YELLOW)⏹ Services stopped$(NC)"

down: ## コンテナを削除
	@docker-compose down
	@echo "$(YELLOW)🗑 Containers removed$(NC)"

restart: ## バックエンドを再起動
	@docker-compose restart backend
	@echo "$(GREEN)♻️ Backend restarted$(NC)"

logs: ## Dockerログを表示
	@docker-compose logs -f

logs-backend: ## バックエンドのログのみ表示
	@docker-compose logs -f backend

logs-db: ## データベースのログのみ表示
	@docker-compose logs -f postgres

db-shell: ## PostgreSQLシェルに接続
	@docker-compose exec postgres psql -U labcareer -d labcareer_db

db-reset: ## データベースをリセット
	@echo "$(RED)Warning: This will delete all data!$(NC)"
	@echo "Press Ctrl+C to cancel, or wait 3 seconds to continue..."
	@sleep 3
	@docker-compose down -v
	@docker-compose up -d
	@echo "$(GREEN)✓ Database reset complete$(NC)"

migration-create: ## マイグレーションファイルを作成
	@read -p "Migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	touch backend/migrations/$${timestamp}_$${name}.up.sql; \
	touch backend/migrations/$${timestamp}_$${name}.down.sql; \
	echo "$(GREEN)✓ Created migration files: $${timestamp}_$${name}$(NC)"

test-backend: ## バックエンドのテストを実行
	@docker-compose exec backend go test ./...

test-frontend: ## フロントエンドのテストを実行
	@cd frontend && npm test

lint: ## リントチェック
	@docker-compose exec backend go vet ./...
	@cd frontend && npm run lint

fmt: ## コードフォーマット
	@docker-compose exec backend go fmt ./...
	@cd frontend && npx prettier --write .

build: ## プロダクションビルド
	@docker-compose exec backend go build -o bin/api cmd/api/main.go
	@cd frontend && npm run build
	@echo "$(GREEN)✓ Production build complete$(NC)"

clean: ## 全データを削除してクリーンアップ
	@echo "$(RED)⚠️  Warning: This will delete everything including volumes!$(NC)"
	@echo "Press Ctrl+C to cancel, or wait 5 seconds to continue..."
	@sleep 5
	@docker-compose down -v
	@rm -rf frontend/node_modules frontend/.next
	@rm -rf backend/tmp backend/bin
	@echo "$(RED)🧹 Cleanup complete$(NC)"

status: ## サービスの状態を確認
	@echo "$(BLUE)Service Status:$(NC)"
	@docker-compose ps
	@echo ""
	@echo "$(BLUE)API Health Check:$(NC)"
	@curl -s http://localhost:8080/health | jq '.' 2>/dev/null || echo "API not responding"
	@echo ""
	@echo "$(BLUE)Database Check:$(NC)"
	@curl -s http://localhost:8080/api/db-check | jq '.' 2>/dev/null || echo "Database check failed"
