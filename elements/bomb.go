package elements

import "encoding/json"

type Bomb struct {
	PosX       int32 `json:"pos_x,omitempty"`
	PoxY       int32 `json:"pox_y,omitempty"`
	BombID     int32 `json:"bomb_id,omitempty"`
	BombRange  int32 `json:"bomb_range,omitempty"`
	BombStatus int32 `json:"bomb_status,omitempty"`
	PlayerID   int32 `json:"player_id,omitempty"`
}

func parseBomb(obj *Obj) (bomb *Bomb, err error) {
	bomb = &Bomb{}
	err = json.Unmarshal(obj.Property, bomb)
	if err != nil {
		return nil, err
	}
	return
}
