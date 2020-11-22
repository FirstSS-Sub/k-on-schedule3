package router

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strings"
)

func jwtAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Firebase SDK のセットアップ
			ctx := context.Background()

			credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_SDK_CREDENTIALS")))
			if err != nil {
				log.Printf("error credentials from json: %v\n", err)
				return err
			}
			opt := option.WithCredentials(credentials)
			// 型アサーションでecho.Context型をcontext.Context型に変換してNewAppの引数に合わせている
			app, err := firebase.NewApp(ctx, nil, opt)
			if err != nil {
				log.Printf("error initializing app: %v", err)
				return err
			}

			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("error getting Firebase client: %n", err)
				return err
			}

			// クライアントから送られてきた JWT 取得
			authHeader := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(authHeader, "Bearer ", "", 1)

			_, err = client.VerifyIDToken(ctx, idToken)
			if err != nil {
				log.Printf("error verifying ID token: %v\n", err)
				c.Response().WriteHeader(http.StatusUnauthorized)
				c.Response().Write([]byte("error verifying ID token\n"))
				return err
			}

			log.Printf("Verifying ID token\n")
			return next(c)
		}
	}
}
