FROM node:16.15.1-alpine as builder

WORKDIR /app

## パッケージをインストール
COPY package*.json .
COPY . .
RUN apk add --no-cache make gcc g++ python3 && \
  npm install && \
  npm rebuild bcrypt --build-from-source && \
  npm run build && \
  apk del make gcc g++ python3



RUN npm run generate

# # 実行用
FROM node:16.15.1-alpine

# WORKDIR /app

## ビルド用のレイヤからコピーする
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json .
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/prisma ./prisma

## Svelteが動く3000ポートを開けておく
EXPOSE 3000

CMD ["node", "./build"]
