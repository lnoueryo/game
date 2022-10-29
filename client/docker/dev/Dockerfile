FROM node:16.15.1-alpine as builder

WORKDIR /app

## パッケージをインストール
COPY package*.json .
COPY playwright.config.ts .
COPY vite.config.ts .
COPY svelte.config.js .
COPY ./.svelte-kit ./.svelte-kit
COPY ./prisma ./prisma
COPY ./.env ./.env

RUN apk add --no-cache make gcc g++ python3 && \
  npm install && \
  npm rebuild bcrypt --build-from-source && \
  apk del make gcc g++ python3

RUN npm run generate

# # 実行用
FROM node:16.15.1-alpine

WORKDIR /app

## ビルド用のレイヤからコピーする
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/playwright.config.ts ./playwright.config.ts
COPY --from=builder /app/vite.config.ts ./vite.config.ts
COPY --from=builder /app/svelte.config.js ./svelte.config.js
COPY --from=builder /app/.svelte-kit ./.svelte-kit
COPY --from=builder /app/.env ./.env
EXPOSE 3000

CMD ["npm", "run", "dev"]
