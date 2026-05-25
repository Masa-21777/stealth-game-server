package models

type GameState struct {
	RoomID        string `json:"room_id"`
	TimeRemaining int    `json:"time_remaining"` // 残り秒数
	Items         bool   `json:"items"`          // アイテム取得済みフラグ
	IsGameOver    bool   `json:"is_game_over"`   // ゲームオーバフラグ
}
