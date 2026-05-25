package models

// イベントの通知用
type GameEvent struct {
	Type      string `json:"type"` // "FOUND" (見つかった), "ITEM_PICKED" (取得)
	PlayerID  string `json:"player_id"`
	TargetID  string `json:"target_id"` // アイテムIDや発見した敵のID
	Timestamp int64  `json:"timestamp"`
}
