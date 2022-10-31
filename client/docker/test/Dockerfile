FROM mcr.microsoft.com/playwright:focal
WORKDIR /app

## パッケージをインストール
COPY ./playwright/package*.json .
COPY ./playwright/playwright.config.ts .
RUN npm install
RUN npx playwright install

EXPOSE 4173

CMD ["npm", "run", "test:watch:docker"]