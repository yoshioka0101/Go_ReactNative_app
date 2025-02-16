package server

import (
	"context"
	"net/http"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
    r := gin.Default()

    // CORS設定
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true, // クッキーや認証情報を許可
    }))

    // ルートの登録
    r.GET("/", s.HelloWorldHandler)
    r.GET("/health", s.healthHandler)

    // Google OAuth 認証開始のルート追加
    r.GET("/auth/google", s.googleAuthHandler)

    // 認証コールバックのルート追加
    r.GET("/auth/:provider/callback", func(c *gin.Context) {
        s.getAuthCallbackFunction(c)
    })

    return r
}


func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getAuthCallbackFunction(c *gin.Context) {
    // プロバイダー名をURLパラメータから取得
    provider := c.Param("provider")

    // コンテキストにプロバイダーを設定
    ctx := context.WithValue(context.Background(), "provider", provider)

    // Gothicを使ってユーザー認証を完了
    user, err := gothic.CompleteUserAuth(c.Writer, c.Request.WithContext(ctx))
    if err != nil {
        // エラーメッセージをレスポンスとして返す
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // ユーザー情報をログに出力
    fmt.Println(user)

    // 認証成功後にリダイレクト
    c.Redirect(http.StatusFound, "http://localhost:5173")
}
