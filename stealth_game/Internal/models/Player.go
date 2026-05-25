package models

// プレイヤーの状態管理
type Player struct {
	ID     string `json:"id"`      // 一意のプレイヤーID
	Name   string `json:"name"`    // プレイヤーの名前
	RoomID string `json:"room_id"` // 所属するルームID

	//座標
	PositionX float64 `json:"position_x"` // X座標
	PositionY float64 `json:"position_y"` // Y座標
	PositionZ float64 `json:"position_z"` // Z座標
	RotationY float64 `json:"rotation_y"` // 左右の向き（Y軸回転）
	RotationX float64 `json:"rotation_x"` // 左右の向き（X軸回転）
	VelocityX float64 `json:"velocity_x"` // 移動速度
	VelocityZ float64 `json:"velocity_z"` // 移動速度

	//フラグ
	IsCrouching bool `json:"is_crouching"` // しゃがみ状態（ステルス状態？現在名称不明）
	IsHidden    bool `json:"is_hidden"`    // 隠れ場所に入っているか(ブッシュみたいなの作成予定)
	IsCaught    bool `json:"is_caught"`    // 捕まったフラグ
}
