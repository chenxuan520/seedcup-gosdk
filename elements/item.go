package elements

import (
	"encoding/json"
)

type ItemType uint8

const (
	ItemNone       ItemType = 0
	ItemBombRange  ItemType = 1
	ItemBombNum    ItemType = 2
	ItemHP         ItemType = 3
	ItemInvincible ItemType = 4
	ItemShield     ItemType = 5
)

type Item struct {
	ItemType ItemType `json:"item_type,omitempty"`
}

func parseItem(obj *Obj) (ItemType, error) {
	item := &Item{}
	err := json.Unmarshal(obj.Property, item)
	return item.ItemType, err
}
