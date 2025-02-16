package  auth

import (
    "log"
    "os"

    "github.com/markbates/goth"
    "github.com/joho/godotenv"
    "github.com/gorilla/sessions"
    "github.com/markbates/goth/gothic"
    "github.com/markbates/goth/providers/google"
)

const (
	key = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
    googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
    redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")

    sessionSecret := os.Getenv("SESSION_SECRET")
    if sessionSecret == "" {
        log.Fatal("SESSION_SECRET is not set in .env")
    }

    store := sessions.NewCookieStore([]byte(sessionSecret))
    store.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   MaxAge,
        HttpOnly: true,
        Secure:   IsProd, // 本番環境ではtrueにする
    }
    gothic.Store = store

    goth.UseProviders(
        google.New(
            googleClientID,
            googleClientSecret,
            redirectURL,
            "email", "profile",
        ),
    )
}