package  auth

import (
    "log"
    "os"

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
		log.Fatal("Error Loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSercret := os.Getenv("GOOGLE_CLIENT_SERCRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store =store
	google.New(googleClientId, googleClientSercret, "http://localhost:3000/auth/google/callback")
}

func min(){

}