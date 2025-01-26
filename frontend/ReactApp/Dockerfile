# ベースイメージとしてnode:18を使用
FROM node:18

# 作業ディレクトリを/appに設定
WORKDIR /app

# 現在のディレクトリの内容をコンテナにコピー
COPY package*.json ./

# npmをインストール
RUN npm install

# 現在のディレクトリの内容をコンテナにコピー
COPY . .

# Expoアプリをトンネルモードで起動
CMD ["npm", "start"]
