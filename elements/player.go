package elements

import "encoding/json"

type Player struct {
	PosX           int32  `json:"pos_x,omitempty"`
	PoxY           int32  `json:"pox_y,omitempty"`
	PlayerID       int32  `json:"player_id,omitempty"`
	PlayerName     string `json:"player_name,omitempty"`
	Alive          bool   `json:"alive,omitempty"`
	HP             int32  `json:"hp,omitempty"`
	ShieldTime     int32  `json:"shield_time,omitempty"`
	InvincibleTime int32  `json:"invincible_time,omitempty"`
	Speed          int32  `json:"speed,omitempty"`
	Score          int32  `json:"score,omitempty"`
	HasGloves      bool   `json:"has_gloves,omitempty"`
	BombRange      int32  `json:"bomb_range,omitempty"`
	BombMaxNum     int32  `json:"bomb_max_num,omitempty"`
	BombNowNum     int32  `json:"bomb_now_num,omitempty"`
}

func parsePlayer(obj *Obj) (player *Player, err error) {
	player = &Player{}
	err = json.Unmarshal(obj.Property, player)
	if err != nil {
		return nil, err
	}
	return
}
