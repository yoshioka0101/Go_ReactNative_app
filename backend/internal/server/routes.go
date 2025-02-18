package server

import (
	"context"
	"net/http"
    "log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
    r := gin.Default()

    // CORS設定
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8081"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true, // クッキーや認証情報を許可
        ExposeHeaders:    []string{"Set-Cookie"},
    }))
    
    // ルートの登録
    r.GET("/", s.HelloWorldHandler)
    r.GET("/health", s.healthHandler)

    // Google OAuth 認証開始のルート追加
    r.GET("/auth/google", s.googleAuthHandler)

    // 
    r.GET("/api/auth/me", s.getUserInfoHandler)

    // ログアウトボタンのルート追加
    r.POST("api/auth/logout", s.logoutHandler)

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
        log.Println("OAuth authentication failed:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth authentication failed", "details": err.Error()})
        return
    }

    // アクセストークンを HttpOnly Cookie に保存
    c.SetCookie("auth_token", user.AccessToken, 3600, "/", "localhost", false, true) 

    // `/dashboard` にリダイレクト（トークンは URL に含めない）
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
