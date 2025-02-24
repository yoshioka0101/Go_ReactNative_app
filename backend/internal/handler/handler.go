package handler

import (
        "context"
        "log"
        "net/http"

        "github.com/gin-gonic/gin"
        "github.com/markbates/goth/gothic"
)

type Database interface {
        Health() map[string]string
}

type Server struct {
        db Database
}

func NewServer(db Database) *Server {
        return &Server{db: db}
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
                log.Println("OAuth authentication failed:", err)
                c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth authentication failed", "details": err.Error()})
                return
        }

        // アクセストークンをHttpOnly Cookieに保存
        c.SetCookie("auth_token", user.AccessToken, 3600, "/", "localhost", false, true)

        // "dashboard"にリダイレク
        c.Redirect(http.StatusFound, "http://localhost:8081/dashboard")
}

func (s *Server) getUserInfoHandler(c *gin.Context) {
        // クライアントから送信された auth_token Cookie を取得
        token, err := c.Cookie("auth_token")
        if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "No auth token found"})
                return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) logoutHandler(c *gin.Context) {
        // logoutHandlerメソッドでCookieを削除する
        c.SetCookie("auth_token", "", -1, "/", "localhost", false, true)
        // 200返す
        c.JSON(http.StatusOK, gin.H{"message": "Logout out successfully"})
}