# Alpine Linuxベースの最新バージョンnodeを採用
FROM node:18-alpine
# コンテナ内の作業ディレクトリ
WORKDIR /app

COPY . .

RUN npm install

EXPOSE 3000

CMD [ "npm", "run", "dev" ]