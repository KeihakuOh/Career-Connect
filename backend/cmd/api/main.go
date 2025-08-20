package main

import (
	"log"
	"os"

	"github.com/KeihakuOh/Career-Connect/internal/config"
	"github.com/KeihakuOh/Career-Connect/internal/database"
	"github.com/KeihakuOh/Career-Connect/internal/server"
)

func main() {
	// 設定を読み込み
	cfg := config.Load()

	// データベース接続
	db, err := database.Connect(cfg)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
		// DBなしでも起動は続行
	} else {
		defer db.Close()
		log.Println("✅ Database connected successfully")
	}

	// サーバー起動
	srv := server.New(cfg, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	if err := srv.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
