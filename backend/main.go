package main

import (
	"20fresh_o/backend/handler"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// 環境変数から値を取得する
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())

	// サーバー起動
	e.Start(":" + port) //ポート番号指定してね
}
