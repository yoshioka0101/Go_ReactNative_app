# ベースイメージとしてgolang:1.23.3を使用
FROM golang:1.23.3

# 作業ディレクトリを/appに設定
WORKDIR /app

# 現在のディレクトリの内容をコンテナにコピー
COPY . .

# ソースコードをビルドする
RUN go build -o backend

# アプリケーションを実行する
CMD ["./backend"]