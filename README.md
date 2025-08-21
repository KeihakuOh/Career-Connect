cat > README.md << 'EOF'

# 🎓 LabCareer - 研究室就活支援プラットフォーム

## 📋 概要

研究室の就活支援を効率化する Web アプリケーション

## 🏗 技術スタック

- **Frontend**: Next.js 14 (App Router), TypeScript, Tailwind CSS
- **Backend**: Go 1.21, Echo Framework
- **Database**: PostgreSQL 15
- **Infrastructure**: Docker, Docker Compose

## 🚀 クイックスタート

### 必要な環境

- Docker & Docker Compose
- Node.js 18+ & npm
- Go 1.21+ (オプション: ローカルでバックエンドを動かす場合)

### セットアップ

```bash
# 1. リポジトリをクローン
git clone https://github.com/yourusername/labcareer.git
cd labcareer

# 2. 初回セットアップ
make setup

# 3. 開発サーバーを起動
make dev

# 4. 別のターミナルでフロントエンドを起動
cd frontend && npm run dev
```
