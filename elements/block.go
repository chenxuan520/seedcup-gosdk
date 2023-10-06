package elements

import "encoding/json"

type Block struct {
	PosX      int32 `json:"pos_x,omitempty"`
	PoxY      int32 `json:"pox_y,omitempty"`
	BlockID   int32 `json:"block_id,omitempty"`
	Removable bool  `json:"removable,omitempty"`
}

func parseBlock(obj *Obj) (block *Block, err error) {
	block = &Block{}
	err = json.Unmarshal(obj.Property, block)
	if err != nil {
		return nil, err
	}
	return
}
