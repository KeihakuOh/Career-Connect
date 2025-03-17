# ベースイメージ
FROM node:18-alpine AS base

# 開発ステージ
FROM base AS development
WORKDIR /app

# 依存関係をインストール
COPY package.json package-lock.json* ./
RUN npm ci

# ポートを公開
EXPOSE 3000

# 開発サーバー起動
CMD ["npm", "run", "dev"]

# ビルドステージ
FROM base AS builder
WORKDIR /app

# 依存関係をコピー
COPY package.json package-lock.json* ./
RUN npm ci

# ソースをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

# 本番ステージ
FROM base AS production
WORKDIR /app

# 環境変数
ENV NODE_ENV production

# パッケージ.jsonとビルド成果物をコピー
COPY --from=builder /app/package.json ./
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules

# ポートを公開
EXPOSE 3000

# サーバー起動
CMD ["npm", "start"]