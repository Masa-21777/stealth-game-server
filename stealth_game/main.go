package main

import (
	"fmt"
	"net/http"
	"stealth_game/Internal/handlers"
)

func main() {
	// 静的ファイルサーバー
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// WebSocket
	http.HandleFunc("/ws", handlers.HandleConnections)

	// ルーム一覧
	http.HandleFunc("/rooms", handlers.HandleRoomList)

	// ランキング（GET: 取得 / POST: スコア追加）
	http.HandleFunc("/ranking", handlers.HandleRanking)

	// ランキングリセット（ターミナルから以下を実行）
	// curl -X DELETE "http://localhost:8080/ranking/reset?key=del123"
	http.HandleFunc("/ranking/reset", handlers.HandleRankingReset)

	fmt.Println("サーバーがポート :8080 で起動しました。接続を待機中...")
	fmt.Println("ランキングリセット: curl -X DELETE \"http://localhost:8080/ranking/reset?key=del123\"")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバー起動エラー:", err)
	}
}
